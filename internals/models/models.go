package models
 
import ( 
 "time" 
) 
 
type User struct { 
 ID        int       `json:"id"` 
 CreatedAt time.Time `json:"created_at"` 
} 
 
type Segment struct { 
 ID        int       `json:"id"` 
 Slug      string    `json:"slug"` 
 CreatedAt time.Time `json:"created_at"` 
} 
 
type UserSegment struct { 
 UserID    int       `json:"user_id"` 
 SegmentID int       `json:"segment_id"` 
 AddedAt   time.Time `json:"added_at"` 
 DeleteTime time.Time `json:"delete_time"` 
}
