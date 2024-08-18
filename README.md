# Instagram-like Image Sharing Platform

This project is a simplified version of an image-sharing platform, similar to Instagram, where users can upload images, add captions, and interact with posts through comments.

## Features

The following user stories have been implemented:

- **Create Posts with Images**: Users can create posts with a single image per post.
- **Set Captions**: Users can add a text caption when creating a post.
- **Comment on Posts**: Users can comment on posts.
- **Delete Comments**: Users can delete their own comments from a post.
- **List Posts with Comments**: Users can retrieve a list of all posts along with the last 2 comments on each post.

## Technology Stack

- **Backend**: Golang with the Gin web framework.
- **Database**: No database implemented due to time-constraint. An in-memory version is used to mimic database interaction
- **Image Processing**: Go's `image` package.

## Installation and Setup

1. **Clone the repository**:
   ```bash
   cd yourproject
   git clone https://github.com/anandh86/instagram.git

2. **Install dependencies**:
   ```bash
   go mod tidy

3. **Run the application**:
   ```bash
   go run .

## API endpoints

The API endpoints are documented using Postman.
https://www.postman.com/universal-crescent-333010/workspace/instagram

## Pending Feature

In the current version, the following features were not implemented due to time constraints:

	1.	Image Resizing: Convert the uploaded images to a 600 x 600 resolution before storing them.
	2.	Post Sorting: Ensure posts are sorted by the number of comments, with the posts having the most comments appearing first.
	3.	Pagination: Implement pagination for the GetAllPosts endpoint to manage large datasets efficiently.
	4.	Image Format Support: BMP format is ignored; only PNG and JPG formats are supported.
	5.	Image Serving Format: Images are served in PNG format instead of JPG.
