# API Endpoints

## 1. Blog List

- Endpoint: /api/blog
- Method: GET
- Headers: -
- Request
  - Parameter: authorize user object
  - Query: -
  - Body: -
- Response:

```
{
    status: Ok/Not Found,
    message: 'Blogs found' / 'Current user does not have blog article yet',
    data:
    {
        [
            {
                id,
                author,
                created_at
                modified_at
                published_at
                slug
                title
                description
                content
                blog_image_url
            },
            .
            .
            .
        ]
    }
}
```

## 2. Single Blog

- Endpoint: /api/blog/:id
- Method: GET
- Headers: -
- Request
  - Parameter: id, authorize user object
  - Query: -
  - Body: -
- Response:

  ```
  {
      status: Ok/Not Found/Forbidden,
      message: 'Blog found' / 'Current user does not have blog article yet' / 'Access forbidden for this user',
      data:
      {
          {
              id,
              author,
              created_at
              modified_at
              published_at
              slug
              title
              description
              content
              blog_image_url
          },
      }
  }
  ```

## 3. Create Blog

- Endpoint: /api/blog/
- Method: POST
- Headers:
  - Auth beared: token
- Request

  - Parameter: authorize user object
  - Query: -
  - Body:
    ```
    {
        title,
        description,
        content,
        image_url,
    }
    ```

- Response:

  ```
  {
      status: Created/Not Authorize/Bad Request,
      message: 'Blog Created'/'Not Authorize'/'Bad request',
      data:
      {
          id,
          slug,
          title,
          description,
          content,
          createdAt,
          updatedAt,
          author: {
              id,
              username,
              fullname,
              role
          }
      }
  }
  ```

## 4. Update Blog

- Endpoint: /api/blog/
- Method: PUT
- Headers:
  - Auth beared: token
- Request

  - Parameter: authorize user object
  - Query: -
  - Body:

  ```
    {
        id,
        title,
        description,
        content,
        image_url,
    }
  ```

- Response:

      {
          status: Ok/Authorize/Not Found/Bad Request/Forbidden,
          message: 'Blog updated'/'Not Authorize'/'Bad request'/'Not Found'/'Forbidde',
          data:
          {
              id,
          }
      }

## 5. Delete Blog

- Endpoint: /api/blog/id
- Method: DELETE
- Headers:
  - Auth beared: token
- Request

  - Parameter: blog id, authorize user object
  - Query: -
  - Body:

- Response:

      {
          status: Ok/Authorize/Not Found/Forbidden,
          message: 'Blog deleted'/'Not Authorize'/'Not Found'/'Forbidde',
          data:
          {
              title,
          }
      }

## 6. User Register

- Endpoint: /api/users/
- Method: POST
- Headers: -
- Request

  - Parameter: -
  - Query: -
  - Body:

    ```
    {
        username,
        email,
        password,
        bio,
        image
    }
    ```

- Response:

  ```
  {
      user:
      {
          username: string;
          image: string;
          email: string;
          bio: string;
          token: string;
      }
  }
  ```

## 7. User Login

- Endpoint: /api/auth/login
- Method: POST
- Headers: Authorization: Bearer token
- Request

  - Parameter: -
  - Query: -
  - Body:

    ```
    {
        username,
        password,
    }
    ```

- Response:

  ```
  {
      status: Ok/Not Found/Bad Request,
      message: ''/'User does not exist'/'Bad request',
      data:
      {
          access_token
      }
  }
  ```

## 8. Find All User

- Endpoint: /api/users
- Method: GET
- Headers: Authorization: Bearer token
- Request
  - Parameter: -
  - Query: -
  - Body:
- Response:

  ```
  {
      users:
      [
          {
              username: string;
              image: string;
              email: string;
              bio: string;
              token: string;
          },
              .
              .
              .
      ]
  }
  ```

## 9. Find One User

- Endpoint: /api/user/id
- Method: GET
- Headers: Authorization: Bearer token
- Request
  - Parameter: user id
  - Query: -
  - Body:
- Response:

  ```
  {
      status: Ok/Not Found/Unauthorized/Forbidden,
      message: 'OK'/'User not found'/'Not Authorized'/'Forbidden Resource',
      data:
      {
          {
              id: string;
              username: string;
              fullname: string;
              image_url: string;
              email: string;
              role: Role;
              createdAt: Date;
          }
      }
  }
  ```

## 10. Delete One User

- Endpoint: /api/user/id
- Method: DELETE
- Headers: Authorization: Bearer token
- Request
  - Parameter: user id
  - Query: -
  - Body:
- Response:

  ```
  {
      status: Ok/Not Found/Unauthorized/Forbidden
      message: 'User deleted'/'User not found'/'Not Authorized'/'Forbidden Resource',
      data:
      {
          username: string;
      }
  }
  ```

<br/>
<br/>
