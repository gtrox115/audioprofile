# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.

## Database Setup

It looks like you chose to set up your application using a database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start it for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	$ buffalo pop create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?



We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)


Start Database:
	$ cd go/src/github.com/gtrox115/profileapi/dbcompose
	$ docker-compose up
	$ docker logs -f my_postgres (See that its working)
	$ docker exec -it my_postgres psql -U postgres -c "create database my_database" (Creates database)

Before starting:
	Update Spotify API key & username in actions/getIds.go

To start up:
	Ensure database is running
	$ cd go/src/github.com/gtrox115/profileapi
	$ buffalo dev
	Visit localhost:3000