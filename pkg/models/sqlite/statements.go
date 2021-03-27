package models

/*------------------------------------------------------*/
/*                                                      */
/*                    POST STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const createPostsTableSQL = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		created TIMESTAMP NOT NULL,
		imageURL TEXT
	);
`
const insertPostSQL = `
	INSERT INTO posts (
		title, content, author, created, imageURL
	) VALUES (?, ?, ?, ?, ?);
`

const getAllPostsSQL = `
	SELECT * FROM posts
`

/*------------------------------------------------------*/
/*                                                      */
/*                    USER STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const createUsersTable = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		login TEXT NOT NULL,
		password TEXT NOT NULL,
		mail TEXT NOT NULL,
		created TIMESTAMP NOT NULL,
		rating INTEGER NOT NULL
	);
`
const insertUserSQL = `
	INSERT INTO users (
		login, password, mail, created, rating
	) VALUES (?, ?, ?, ?, ?);
`
