package time

import (
	"fmt"
	"time"
)

// TimestampToSrt — returns timestamp in srt format
func TimestampToSrt(t time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d", t/time.Hour, (t%time.Hour)/time.Minute, (t%time.Minute)/time.Second, (t%time.Second)/time.Millisecond)
}
