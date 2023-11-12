# Steps to api

Golang + Gin

## what is the purpose?

    as a gamer
    i would like to play games with my friends
    to play with my friends
    i would like to see what games my friends have

    I would like to be able to login with and connect my game platform accounts

    when using an api tool like insomnia
    i would like to access data showing 
        user
            id
            username
            email
            password
            ownedGames [game]
            friends [user]
            (might need to store auth tokens)

        game
            id
            title
            platforms [platform]
            prices [gamePlatformPrices]

        platform
            id
            name
            games [gamePlatformPrices]

        gamePlatformPrices
            id
            gameId
                one to many association
            platformId
                many to one association
            crossplayEnabled bool
            price
            

## setting up connections

### db to use?

    research reasons for using different dbs

        mongodb = non-relational
        mysql = relational

        PostgreSQL = mysql but less bad = the one im using

        a lot of my data is relational so i think it would be a good
        idea to use a relational database

### JSON format from database?

    ask aurora, re the JSON format thing he mentioned was in his airplane game back end
        func to upsert to JSON?

### Endpoints

    user
        POST signup
            
        POST login
        DELETE removeUser
        GET me
        GET singleUser
        GET manyUsers
            params: x many users

        POST addGame
        DELETE removeGame

        UPDATE addConnection
        UPDATE changeUserInfo
        UPDATE changePassword

        friends
            POST newFriend
            DELETE removeFriend

    game
        GET singleGame
        GET manyGames
            params: x many games

    platform
        GET platform

    search
        POST search
            text
            tags
            platform
            priceRange [low, high]

            return games = [game]

### auth middleware

    Auth 
        pull user data from other platforms on connection?
            routing responsibility -> redirect to platform for auth
                send data as headers for create user? 

### Database

    user
        id
            required true
        username
            required true
        email
            required true
        password
            required true
        ownedGames [game]
            required false
        friends [user]
            required false

        on creation -> hash password

    game
        id
        title
        platforms [platform]
        prices [gamePlatformPrices]

    platform
        id
        name

    gamePlatformPrices
        id
        gameId
        platformId
        crossplayEnabled bool
        price

### sample data

    games
        fill from API call to platforms?

            upsert func -> convert platform
            data to local data

            func to grab all games from platform
                if already games in database
                    func to search database for 
                    matching titles and add 
                    platform and crossplay status
                else    
                    add to database

    because auth will need to upsert data to database

        early dev = before auth finalized
            use fake users
        late dev = after auth = front end login/signin complete
            use real user from dev team (me lol) 

## fill in the lines

### CRUD ops

    depends on db and interface package?

    PostgreSQL

### Routing

### etc...

## testing

### confirm working

### unit testing

### fill out README

## ideas for future dev

    recommendation algorithm for games and friends

    front-end

    forum
    official events/mixers