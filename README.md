# library-system

This is CRUD web application to manage books in library.
You can add a book.
![image](https://github.com/user-attachments/assets/83353438-c4e7-45a7-9252-721b4364e4c6)

You can edit an existing book.
![image](https://github.com/user-attachments/assets/29996370-426c-4583-a83f-8259293d124f)

You can see the list of the books and delete a book.
![image](https://github.com/user-attachments/assets/0f7db12d-438a-49a4-887f-2278c9987554)

You can also search a book by book name or author.
![image](https://github.com/user-attachments/assets/00cb4ed2-e8e5-491a-bf64-6fad11cadf3e)

[Sqlx](https://github.com/jmoiron/sqlx) has been used to connect with Database. I used PostgreSQL as the database driver.
You can modify the below code in the main.go file to use another database.
```
db, err := sqlx.Connect("postgres", "user=postgres password=Anubis0912 dbname=library sslmode=disable")
```
