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

const createUsersTableSQL = `
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

/*------------------------------------------------------*/
/*                                                      */
/*                   COMMENT STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const createCommentsTableSQL = `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postID INTEGER NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		created TIMESTAMP NOT NULL,
	);
`

const insertCommentSQL = `
	INSERT INTO comments (
		postID, author, content, created
	) VALUES (?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   CATEGORY STATEMENTS                */
/*                                                      */
/*------------------------------------------------------*/

const createCategorySQL = `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
`

const createCategoryPostLinkSQL = `
	CREATE TABLE IF NOT EXISTS categoryPostLink (
		postID INTEGER NOT NULL,
		name STRING NOT NULL
	);
`

const insertCategoriesSQL = `
	INSERT INTO categories (
		name
	) VALUES (?);
`

const insertCategoryPostLink = `
	INSERT INTO categoryPostLink (
		postID, name
	) VALUES (?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                    RATING STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/
