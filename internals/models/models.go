package models
 
import ( 
 "time" 
) 
 
type User struct { 
 ID       int64       `json:"id"` 
 CreatedAt time.Time `json:"created_at,omitempty"` 
} 
 
type AddRequest struct{
	ID int64  `json:"id"` 
	AddSegment[] string `json:"add_segment,omitempty"` 
	DeleteSegment[] string  `json:"delete_segment,omitempty"` 
	DeleteTime time.Time `json:"delete_time,omitempty"` 
}
type Segment struct { 
 ID        int       `json:"id,omitempty"` 
 Slug      string    `json:"slug"` 
 Procent   int    	`json:"procent,omitempty"` 
 CreatedAt time.Time `json:"created_at,omitempty"` 
} 
 
type UserSegment struct { 
 UserID    int       `json:"user_id"` 
 SegmentID int       `json:"segment_id"` 
 AddedAt   time.Time `json:"added_at"` 
 DeleteTime time.Time `json:"delete_time,omitempty"` 
}
