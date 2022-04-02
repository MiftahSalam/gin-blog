<br/>
<br/>

# Database

## 1. Tabel List

- blog
  - columns
    - id: int
    - author (Foreign key): int
    - 
    - 
- user
  - columns
    - id: int
    - username: varchar
    - email: varchar
    - password: varchar
    - bio: varchar
    - image: varchar
- follow
  - columns
    - id: int
    - created_at: timestamp
    - updated_at: timestamp
    - deleted_at: timestamp
    - following_id: int
    - followed_by_id: int
