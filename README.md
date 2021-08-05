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
GET localhost:8080/question/number2/list?search=batman&page=1
```

```console
Response:
{
    "success": true,
    "code": 200,
    "message": "Success",
    "data": {
        "Title": "Batman: Under the Red Hood",
        "Year": "2010",
        "Rated": "PG-13",
        "Released": "27 Jul 2010",
        "Runtime": "75 min",
        "Genre": "Animation, Action, Crime, Drama, Mystery, Sci-Fi, Thriller",
        "Director": "Brandon Vietti",
        "Writer": "Judd Winick, Bob Kane (Batman created by)",
        "Actors": "Bruce Greenwood, Jensen Ackles, John DiMaggio, Neil Patrick Harris",
        "Plot": "There's a mystery afoot in Gotham City, and Batman must go toe-to-toe with a mysterious vigilante, who goes by the name of Red Hood. Subsequently, old wounds reopen and old, once buried memories come into the light.",
        "Language": "English",
        "Country": "USA",
        "Awards": "1 nomination.",
        "Poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
        "Ratings": [
            {
                "Source": "Internet Movie Database",
                "Value": "8.1/10"
            },
            {
                "Source": "Rotten Tomatoes",
                "Value": "100%"
            }
        ],
        "Metascore": "N/A",
        "imdbRating": "8.1",
        "imdbVotes": "56,401",
        "imdbID": "tt1569923",
        "Type": "movie",
        "DVD": "N/A",
        "BoxOffice": "N/A",
        "Production": "DC Entertainment, Warner Bros. Animation",
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