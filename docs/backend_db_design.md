<br/>
<br/>

# Database

## 1. Tabel List

- article_models
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - slug
    - title
    - description
    - body
    - author_id (Foreign key): int

- article_user_model
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - user_id (Foreign key): int

- comment_models
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - article_id (Foreign key): int
    - author_id (Foreign key): int
    - body

- favoite_models
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - favorite_id (Foreign key): int
    - favorite_by_id (Foreign key): int

- tag_models
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - tag

- article_tags
  - columns
    - tag_model_id (Foreign key): int
    - article_model_id (Foreign key): int

- users
  - columns
    - id: int
    - username 
    - email
    - bio
    - image_url
    - password

- follow
  - columns
    - id: int
    - created_at
    - updated_at
    - deleted_at
    - following_id: int
    - followed_by_id: int
