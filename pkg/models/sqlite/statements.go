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
		created_at TIMESTAMP NOT NULL,
		image_url TEXT
	);
`
const insertPostSQL = `
	INSERT INTO posts (
		title, content, author, created_at, image_url
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
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		rating INTEGER NOT NULL
	);
`
const insertUserSQL = `
	INSERT INTO users (
		username, password, email, created_at, rating
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
		post_id INTEGER NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
	);
`

const insertCommentSQL = `
	INSERT INTO comments (
		post_id, author, content, created_at
	) VALUES (?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   CATEGORY STATEMENTS                */
/*                                                      */
/*------------------------------------------------------*/

const createCategoryTableSQL = `
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

const insertCategoryPostLinkSQL = `
	INSERT INTO categoryPostLink (
		postID, name
	) VALUES (?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                    RATING STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const createRatingPostSQL = `
	CREATE TABLE IF NOT EXISTS ratingPosts (
		postID INTEGER NOT NULL,
		author STRING NOT NULL,
		value INTEGER NOT NULL
	);
`

const createRatingCommentSQL = `
	CREATE TABLE IF NOT EXISTS ratingComments (
		commentID INTEGER NOT NULL,
		author STRING NOT NULL,
		value INTEGER NOT NULL
	);
`

const insertRatingPostSQL = `
	INSERT INTO ratingPosts (
		postID, author, value
	) VALUES (?, ?, ?);
`

const insertRatingCommentSQL = `
	INSERT INTO ratingComments (
		commentID, author, value
	) VALUES (?, ?, ?);
`

