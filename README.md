# WaysGallery

WaysGallery is a place where designer creators gather and share. This is a demo website that can be used to understand more about fullstack website development.

This website was created using [Bootstrap](https://getbootstrap.com/) and [TailwindCSS](https://tailwindcss.com/) for styling, [React](https://react.dev/) for the frontend framework, [GO](https://go.dev/) for the backend with the [Echo](https://echo.labstack.com/) framework, [Gorm](https://gorm.io/) for querying the database, [PostgreSQL](https://postgresql.org/) as the database, and [Cloudinary](https://cloudinary.com/) as as the file storage.

## Available Features

- Modal Login & Register.
- Landing Pages.
- A page showing all posts.
- Feature to follow other creators.
- Filter to see only posts from creators that are followed on the all post display page.
- Search feature to find posts on the all post display page.
- Profile page for each creator.
- Job list feature that you want to show off on each creator's profile.
- Features of hiring other creators.
- Page to manage the process of hiring other creators.
- API Endpoint for all above features.
- Password Hashing Middleware for each User using [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
- Middleware Upload File to upload files from user input files.
- Middleware to authenticate by creating a Token from [JWT](https://jwt.io/).
- Payment Gateways using [Midtrans](midtrans.com).
