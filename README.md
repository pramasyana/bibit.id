#### Question Number 1
```console
Table:
id | username | parent
1  | Ali      | 2
2  | Budi     | 0
3  | Cecep    | 1

Expected:
id | username | parentUsername
1  | Ali      | Budi
2  | Budi     | null
3  | Cecep    | Ali
```

```console
GET localhost:8080/question/number1
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "database": "Mysql",
        "query": "SELECT id, username, (select username from `user` where id=us.parent) as parentUsername FROM `user` as us"
    }
}
```



#### Question Number 2 List

```console
GET localhost:8080/question/number2/list?search=batman&page=1
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "Search": [
            {
                "Title": "Batman Begins",
                "Year": "2005",
                "imdbID": "tt0372784",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman v Superman: Dawn of Justice",
                "Year": "2016",
                "imdbID": "tt2975590",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman",
                "Year": "1989",
                "imdbID": "tt0096895",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"
            },
            {
                "Title": "Batman Returns",
                "Year": "1992",
                "imdbID": "tt0103776",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"
            },
            {
                "Title": "Batman Forever",
                "Year": "1995",
                "imdbID": "tt0112462",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"
            },
            {
                "Title": "Batman & Robin",
                "Year": "1997",
                "imdbID": "tt0118688",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"
            },
            {
                "Title": "The Lego Batman Movie",
                "Year": "2017",
                "imdbID": "tt4116284",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: The Animated Series",
                "Year": "1992–1995",
                "imdbID": "tt0103359",
                "Type": "series",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: Under the Red Hood",
                "Year": "2010",
                "imdbID": "tt1569923",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: The Dark Knight Returns, Part 1",
                "Year": "2012",
                "imdbID": "tt2313197",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"
            }
        ],
        "Response": "True"
    }
}
```



#### Question Number 2 Detail

```console
GET localhost:8080/question/number2/detail?id=tt2820852
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "Title": "Fast & Furious 7",
        "Year": "2015",
        "Rated": "PG-13",
        "Released": "03 Apr 2015",
        "Runtime": "137 min",
        "Genre": "Action, Adventure, Thriller",
        "Director": "James Wan",
        "Writer": "Chris Morgan, Gary Scott Thompson",
        "Actors": "Vin Diesel, Paul Walker, Dwayne Johnson",
        "Plot": "Deckard Shaw seeks revenge against Dominic Toretto and his family for his comatose brother.",
        "Language": "English, Thai, Arabic, Spanish",
        "Country": "United States, China, Japan, Canada, United Arab Emirates",
        "Awards": "35 wins & 36 nominations",
        "Poster": "https://m.media-amazon.com/images/M/MV5BMTQxOTA2NDUzOV5BMl5BanBnXkFtZTgwNzY2MTMxMzE@._V1_SX300.jpg",
        "Ratings": [
            {
                "Source": "Internet Movie Database",
                "Value": "7.1/10"
            },
            {
                "Source": "Rotten Tomatoes",
                "Value": "82%"
            },
            {
                "Source": "Metacritic",
                "Value": "67/100"
            }
        ],
        "Metascore": "67",
        "imdbRating": "7.1",
        "imdbVotes": "369,515",
        "imdbID": "tt2820852",
        "Type": "movie",
        "DVD": "01 Mar 2016",
        "BoxOffice": "$353,007,020",
        "Production": "Original Film",
        "Website": "N/A",
        "Response": "True"
    }
}
```



#### Question Number 3

```console
POST localhost:8080/question/number3
```

```console
Body:
{
    "text": "ini adalah (Ryan Pramasyana)"
}
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "text": "ini adalah (Ryan Pramasyana)",
        "result": "Ryan Pramasyana"
    }
}
```



#### Question Number 4

```console
POST localhost:8080/question/number4
```

```console
Body:
{
    "text": ["kita", "atik", "tika", "aku", "kia", "makan", "kua"]
}
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "text": [
            "kita",
            "atik",
            "tika",
            "aku",
            "kia",
            "makan",
            "kua"
        ],
        "result": [
            [
                "kita",
                "atik",
                "tika"
            ],
            [
                "aku",
                "kua"
            ],
            [
                "kia"
            ],
            [
                "makan"
            ]
        ]
    }
}
```