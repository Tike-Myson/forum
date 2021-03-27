package models

/*------------------------------------------------------*/
/*                                                      */
/*                    POST STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const CreatePostsTableSQL = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		image_url TEXT
	);
`
const InsertPostSQL = `
	INSERT INTO posts (
		title, content, author, created_at, image_url
	) VALUES (?, ?, ?, ?, ?);
`

const GetAllPostsSQL = `
	SELECT * FROM posts
`

/*------------------------------------------------------*/
/*                                                      */
/*                    USER STATEMENTS                   */
/*                                                      */
/*------------------------------------------------------*/

const CreateUsersTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		rating INTEGER NOT NULL
	);
`
const InsertUserSQL = `
	INSERT INTO users (
		username, password, email, created_at, rating
	) VALUES (?, ?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   COMMENT STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const CreateCommentsTableSQL = `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
	);
`

const InsertCommentSQL = `
	INSERT INTO comments (
		post_id, author, content, created_at
	) VALUES (?, ?, ?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                   CATEGORY STATEMENTS                */
/*                                                      */
/*------------------------------------------------------*/

const CreateCategoryTableSQL = `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
`

const CreateCategoryPostLinkSQL = `
	CREATE TABLE IF NOT EXISTS categoryPostLink (
		postID INTEGER NOT NULL,
		name STRING NOT NULL
	);
`

const InsertCategoriesSQL = `
	INSERT INTO categories (
		name
	) VALUES (?);
`

const InsertCategoryPostLinkSQL = `
	INSERT INTO categoryPostLink (
		postID, name
	) VALUES (?, ?);
`

/*------------------------------------------------------*/
/*                                                      */
/*                    RATING STATEMENTS                 */
/*                                                      */
/*------------------------------------------------------*/

const CreateRatingPostSQL = `
	CREATE TABLE IF NOT EXISTS ratingPosts (
		postID INTEGER NOT NULL,
		author STRING NOT NULL,
		value INTEGER NOT NULL
	);
`

const CreateRatingCommentSQL = `
	CREATE TABLE IF NOT EXISTS ratingComments (
		commentID INTEGER NOT NULL,
		author STRING NOT NULL,
		value INTEGER NOT NULL
	);
`

const InsertRatingPostSQL = `
	INSERT INTO ratingPosts (
		postID, author, value
	) VALUES (?, ?, ?);
`

const InsertRatingCommentSQL = `
	INSERT INTO ratingComments (
		commentID, author, value
	) VALUES (?, ?, ?);
`

