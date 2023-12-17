package sessions

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/mtstnt/urunan/database"
)

type sessionMap map[string]int64
type expireTrackerMap map[int64]string

const tickInterval = 5 * time.Second

var (
	sessions      sessionMap
	expireTracker expireTrackerMap
	mu            sync.Mutex
)

func Initialize() {
	sessions = make(sessionMap)
	expireTracker = make(expireTrackerMap)

	go func(t *time.Ticker) {
		cur := <-t.C
		slog.Debug("Checking expire tracker.")
		mu.Lock()
		for k, v := range expireTracker {
			if cur.Unix() > k {
				slog.Debug("Token " + v + " expired.")
				delete(expireTracker, k)
				delete(sessions, v)
			}
		}
		mu.Unlock()
	}(time.NewTicker(tickInterval))
}

func Get(token string, ctx context.Context, db database.Querier) (database.User, error) {
	userID, isExist := sessions[token]
	if !isExist {
		return database.User{}, fmt.Errorf("not found")
	}
	return db.GetUserByID(ctx, userID)
}

func Register(token string, expiresAt int64, userID int64) {
	sessions[token] = userID
	expireTracker[expiresAt] = token
}

func Revoke(token string) {
	delete(sessions, token)
	for k, v := range expireTracker {
		if v == token {
			delete(expireTracker, k)
			break
		}
	}
}

func Exists(token string) bool {
	_, isExist := sessions[token]
	return isExist
}
