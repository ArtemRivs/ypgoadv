package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(dataBaseDsn string) *Storage {
	ls := Storage{}
	var err error
	log.Println("dataBaseDsn:", dataBaseDsn)
	ls.db, err = sql.Open("postgres", dataBaseDsn)
	if err != nil {
		panic(err)
	}
	err = ls.db.Ping()
	if err != nil {
		panic(err)
	}
	// _, err = ls.db.Exec("CREATE TABLE IF NOT EXISTS links (short_code VARCHAR(8), origin_url TEXT UNIQUE, user_id VARCHAR(16), is_deleted BOOLEAN DEFAULT FALSE);")
	_, err = ls.db.Exec("CREATE TABLE IF NOT EXISTS links (short_code VARCHAR NOT NULL, origin_url TEXT UNIQUE, user_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE);")

	if err != nil {
		panic(err)
	}
	return &ls
}

func (ls *Storage) SaveOriginLink(shortCode, longURL, userID string) (string, error) {
	_, err := ls.db.Exec("INSERT INTO links (short_code, origin_url, user_id) VALUES ($1, $2, $3);", shortCode, longURL, userID)
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == pgerrcode.UniqueViolation {
		ls.db.QueryRow("SELECT short_code FROM links WHERE origin_url = ($1);", longURL).Scan(&shortCode)
		return shortCode, fmt.Errorf(`%w`, ErrExistedURL)
	}
	return shortCode, err
}

func (ls *Storage) GetOriginLink(shortCode string) (string, error) {
	var originURL string
	var isDeleted bool
	row := ls.db.QueryRow("SELECT origin_url, is_deleted FROM links WHERE short_code = ($1);", shortCode)
	err := row.Scan(&originURL, &isDeleted)

	if err != nil {
		return "", err
	}
	if isDeleted {
		return "", fmt.Errorf(`%w`, ErrDeletedURL)
	}

	return originURL, nil
}

func (ls *Storage) GetUserOriginLinks(userID string) (map[string]string, error) {
	ul := make(map[string]string)
	rows, err := ls.db.Query("SELECT origin_url, short_code FROM links WHERE user_id = ($1) AND is_deleted=false;", userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var originURL, shortCode string
		err = rows.Scan(&originURL, &shortCode)
		if err != nil {
			return nil, err
		}
		ul[shortCode] = originURL
	}
	err = rows.Err()
	if err != nil {
		return ul, nil
	}
	return ul, nil
}

// func (ls *Storage) DeleteUserLinks(userID string, shortCodes []string) {

// 	log.Println("shortCodes:", shortCodes)

// 	tx, err := ls.db.Begin()
// 	if err != nil {
// 		log.Printf("URL delete error: %v\n", err)
// 		return
// 	}
// 	q := "UPDATE links SET is_deleted = true WHERE short_code = $1 AND user_id = $2;"
// 	stmt, err := ls.db.Prepare(q)
// 	if err != nil {
// 		log.Printf("DB prepare querry error: %v\n", err)
// 		return
// 	}
// 	defer stmt.Close()

// 	for _, r := range shortCodes {
// 		_, err = stmt.Exec(r, userID)
// 		if err != nil {
// 			err = tx.Rollback()
// 			if err != nil {
// 				log.Printf("URL delete error: %v\n", err)
// 				return
// 			}
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Printf("commit error: %v\n", err)
// 		return
// 	}
// }

func (ls *Storage) Close() {
	ls.db.Close()
}

// func (ls *Storage) DeleteUserLinks(ctx context.Context, userID string, shortCodes []string) error
func (ls *Storage) DeleteUserLinks(ctx context.Context, userID string, shortCodes []string) error {
	batchSize := 100 // set the batch size to 100 for example purposes
	codeChunks := chunkStringSlice(shortCodes, batchSize)
	errChan := make(chan error, len(codeChunks))

	for _, chunk := range codeChunks {
		go func(c []string) {
			query := "UPDATE links SET is_deleted = true WHERE user_id = $1 AND short_code = ANY($2)"
			_, err := ls.db.Exec(query, userID, pq.Array(c))
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
		}(chunk)
	}

	// wait for all goroutines to finish
	for i := 0; i < len(codeChunks); i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

// helper function to split a string slice into smaller chunks
func chunkStringSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// func fanOut(input []string, n int) []chan string {
// 	chs := make([]chan string, 0, n)
// 	for i, val := range input {
// 		ch := make(chan string, 1)
// 		ch <- val
// 		chs = append(chs, ch)
// 		close(chs[i])
// 	}
// 	return chs
// }

// func newWorker(ctx context.Context, stmt *sql.Stmt, tx *sql.Tx, userID string, inputCh <-chan string, errCh chan<- error, wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func() {
// 		var defErr error
// 		defer func() {
// 			if defErr != nil {
// 				select {
// 				case errCh <- defErr:
// 				case <-ctx.Done():
// 					log.Println("cancel delete")
// 				}
// 			}
// 			wg.Done()
// 		}()
// 		for id := range inputCh {
// 			if _, err := stmt.ExecContext(ctx, id, userID); err != nil {
// 				if err = tx.Rollback(); err != nil {
// 					defErr = err
// 					return
// 				}
// 				defErr = err
// 				return
// 			}
// 		}
// 	}()
// }

// func (ls *Storage) DeleteUserLinks(ctx context.Context, userID string, shortCodes []string) error {

// 	// log.Println("userID:", userID)
// 	// log.Println("shortCodes:", shortCodes)

// 	n := len(shortCodes)

// 	tx, err := ls.db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	if len(shortCodes) == 0 {
// 		return errors.New("the list of URLs is empty")
// 	}

// 	fanOutChs := fanOut(shortCodes, n)
// 	stmt, err := tx.PrepareContext(ctx, "UPDATE links SET is_deleted=TRUE WHERE short_code=$1 AND user_id=$2")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	wg := &sync.WaitGroup{}
// 	errCh := make(chan error)
// 	for _, item := range fanOutChs {
// 		newWorker(ctx, stmt, tx, userID, item, errCh, wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(errCh)
// 	}()

// 	if err = <-errCh; err != nil {
// 		log.Println(err)
// 		//cancel()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	log.Println("end userID:", userID)
// 	log.Println("end shortCodes:", shortCodes)

// 	return nil
// }
