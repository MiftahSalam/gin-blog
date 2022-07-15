# API Endpoints

## 1. Article List

- Endpoint: /api/article/
- Method: GET
- Headers: -
- Request
  - Parameter: -
  - Query:
    - tag : string
    - author : string
    - favorite by : string
    - limit : integer
    - offset : integer
  - Body: -
- Response:
  - http status : Ok/NotFound
  - Body :

```
{
    articles :
    [
      {
        id,
        title,
        slug,
        description,
        body,
        created_at,
        updated_at,
        author: {
          id,
          username,
          bio,
          image_url,
          following
        },
        tags,
        favorited,
        favoritedCount
      },
    ],
    articlesCount,
}
```

## 3. Article feeds

- Endpoint: /api/article/feed
- Method: GET
- Headers: Auth Bearer token
- Request
  - Parameter: -
  - Query:
    - limit : integer
    - offset : integer
  - Body: -
- Response:
  - http status : Ok/NotFound/Unauthorized
  - Body :

```
{
    articles :
    [
      {
        id,
        title,
        slug,
        description,
        body,
        created_at,
        updated_at,
        author: {
          id,
          username,
          bio,
          image_url,
          following
        },
        tags,
        favorited,
        favoritedCount
      },
    ],
    articlesCount,
}
```

## 3. Single article

- Endpoint: /api/article/:slug
- Method: GET
- Headers: Auth Bearer token
- Request
  - Parameter: article slug
  - Query: -
  - Body: -
- Response:
  - http status : Ok/NotFound
  - Body :

```
{
    article :
    {
      id,
      title,
      slug,
      description,
      body,
      created_at,
      updated_at,
      author: {
        id,
        username,
        bio,
        image_url,
        following
      },
      tags,
      favorited,
      favoritedCount
    },
}
```

## 4. Create article

- Endpoint: /api/article/
- Method: POST
- Headers:
  - Auth beared: token
  - Content Type : application/json
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
      article: {
        title, -> required
        description,
        body,
        tags,
      }
    }
    ```
- Response:
  - http status: Created/Unauthorized/Bad Request
  - body :
    ```
    {
        article :
        {
          id,
          title,
          slug,
          description,
          body,
          created_at,
          updated_at,
          author: {
            id,
            username,
            bio,
            image_url,
            following
          },
          tags,
          favorited,
          favoritedCount
        },
      }
    ```

## 5. Update article

- Endpoint: /api/article/:slug
- Method: PUT
- Headers:
  - Auth beared: token
  - Content Type : application/json
- Request
  - Parameter: slug
  - Query: -
  - Body:
    ```
    {
        title, -> required
        description,
        body,
        tags,
    }
    ```
- Response:
  - http status: Ok/Unauthorized/Bad Request/Not Found
  - body :
    ```
    {
        article :
        {
          id,
          title,
          slug,
          description,
          body,
          created_at,
          updated_at,
          author: {
            id,
            username,
            bio,
            image_url,
            following
          },
          tags,
          favorited,
          favoritedCount
        },
      }
    ```

## 6. Favorite article

- Endpoint: /api/article/:slug/favorite
- Method: POST
- Headers:
  - Auth bearer: token
- Request
  - Parameter: slug
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized/Bad Request/Not Found
  - body :
    ```
    {
        article :
        {
          id,
          title,
          slug,
          description,
          body,
          created_at,
          updated_at,
          author: {
            id,
            username,
            bio,
            image_url,
            following
          },
          tags,
          favorited,
          favoritedCount
        },
      }
    ```

## 7. UnFavorite article

- Endpoint: /api/article/:slug/favorite
- Method: DELETE
- Headers:
  - Auth bearer: token
- Request
  - Parameter: slug
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized/Bad Request/Not Found
  - body :
    ```
    {
        article :
        {
          id,
          title,
          slug,
          description,
          body,
          created_at,
          updated_at,
          author: {
            id,
            username,
            bio,
            image_url,
            following
          },
          tags,
          favorited,
          favoritedCount
        },
      }
    ```

## 8. Create article comment

- Endpoint: /api/article/:slug/comment
- Method: POST
- Headers:
  - Auth beared: token
  - Content Type : application/json
- Request
  - Parameter: slug
  - Query: -
  - Body:
    ```
    {
        body,
    }
    ```
- Response:
  - http status: Created/Unauthorized/Bad Request/NotFound
  - body :
    ```
    {
        comment :
        {
          id,
          body,
          created_at,
          updated_at,
          author: {
            id,
            username,
            bio,
            image_url,
            following
          },
        },
      }
    ```

## 9. List article comments

- Endpoint: /api/article/:slug/comment
- Method: GET
- Headers:
  - Auth beared: token
- Request
  - Parameter: slug
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized/NotFound
  - body :
    ```
    {
        comments :
        [
          {
            id,
            body,
            created_at,
            updated_at,
            author: {
              id,
              username,
              bio,
              image_url,
              following
            },
          },
        ]
    }
    ```

## 10. Delete article comment

- Endpoint: /api/article/:slug/comment
- Method: DELETE
- Headers:
  - Auth beared: token
- Request
  - Parameter: slug
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized/NotFound
  - body :
    ```
    {
        comment : Deleted
    }
    ```

## 11. Delete article

- Endpoint: /api/article/:slug
- Method: DELETE
- Headers:
  - Auth beared: token
- Request
  - Parameter: slug
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized/NotFound
  - body :
    ```
    {
        article : Deleted
    }
    ```

## 12. Create user (register)

- Endpoint: /api/users/
- Method: POST
- Headers:
  - Content Type : application/json
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
        username, -> required
        email, -> required
        password, -> required
        bio,
        image_url,
    }
    ```
- Response:
  - http status: Created/Bad Request
  - body :
    ```
    {
      user : {
        username,
        email,
        bio,
        image_url,
        token,
      }
    }
    ```

## 13. login user

- Endpoint: /api/users/login
- Method: POST
- Headers:
  - Content Type : application/json
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
        user : {
          email, -> required
          password, -> required
        }
    }
    ```
- Response:
  - http status: Ok/Bad Request
  - body :
    ```
    {
      user : {
        username,
        email,
        bio,
        image_url,
        token,
      }
    }
    ```

## 14. List users

- Endpoint: /api/users/
- Method: GET
- Headers:
  - Auth bearer: token
- Request
  - Parameter: -
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Unauthorized
  - body :
    ```
    {
        users :
        [
          {
            username,
            email,
            bio,
            image_url,
            token,
          },
        ]
    }
    ```

## 15. Update user

- Endpoint: /api/users/:id
- Method: PUT
- Headers:
  - Content Type : application/json
  - Auth bearer: token
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
        username, -> required
        email, -> required
        password, -> required
        bio,
        image_url,
    }
    ```
- Response:
  - http status: Ok/Bad Request/Unauthorized/Not Found
  - body :
    ```
    {
      user : {
        username,
        email,
        bio,
        image_url,
        token,
      }
    }
    ```

## 16. Get user profile

- Endpoint: /api/users/:username
- Method: GET
- Headers: -
- Request
  - Parameter: username
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Not Found
  - body :
    ```
    {
      user : {
        id,
        username,
        email,
        bio,
        image_url,
        following,
      }
    }
    ```

## 17. Follow user

- Endpoint: /api/users/:username/follow
- Method: POST
- Headers:
  - Auth bearer: token
- Request
  - Parameter: username
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Not Found
  - body :
    ```
    {
      user : {
        id,
        username,
        email,
        bio,
        image_url,
        following,
      }
    }
    ```

## 18. Get user following

- Endpoint: /api/users/following
- Method: GET
- Headers:
  - Auth bearer: token
- Request
  - Parameter: -
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Not Found
  - body :
    ```
    {
      [
        {
          id,
          username,
          email,
          bio,
          image_url,
          following,
        },
      ]
    }
    ```

## 19. User unfollow

- Endpoint: /api/users/:username/follow
- Method: DELETE
- Headers:
  - Auth bearer: token
- Request
  - Parameter: -
  - Query: -
  - Body: -
- Response:
  - http status: Ok/Not Found
  - body :
    ` { user : { id, username, email, bio, image_url, following, } } `
    <br/>
    <br/>
