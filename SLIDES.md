---
author: Jordan Brauer
date: MMMM dd, YYYY
---

# (^０^)ノHello!

My name is Jordan Brauer

I've been doing software for a few years both professionally, and as a hobby.

1. Pepper (Tipping Canoe)
2. Gubagoo
3. Neo Financial

- _GitHub_ as [@jordanbrauer](https://github.com/jordanbrauer)
- _Twitter_ as [@jorbandrauer](https://twitter.com/jorbandrauer)
- _Instagram_ as [@jorbandrauer](https://instagram.com/jorbandrauer)

_Alright, moving on_ ...

---

## __〆(￣ー￣ ) Agenda

We will be visiting the following topics during this talk

1. Myth Busting
2. A Brief History
3. Crash Course
4. Demonstration
5. Questions _(please save until the end)_

---

# ヽ(°〇°)ﾉ Misconceptions

1. It's a fad
2. It requires a dedicated server
3. The client is querying the database
4. It's expensive to adopt and rewrite for
5. Trying to replace REST/SOAP

---

# (￢_￢) Who Made GraphQL?

_Lee Byron_ at **Facebook**; Started as an **internal tool in 2012** which was
then **open-sourced in 2015**.

In November of 2018, GraphQL was moved from Facebook to the new GraphQL
Foundation, hosted by the non-profit Linux Foundation.

- On track for Lee's original adoption timeline

---

...

_It's not a fad._

---

# (・_・ヾ What is GraphQL?

An open-source query language & type system.

##### Schema Definition Language (Server)

```graphql
type Query {
    ping: String!
}

type Mutation {
    sum(numbers: [Int!]!): Int!
}
```

##### Query Language (Client)

```graphql
query PingPong {
    ping
}

mutation Add($left: Int! $right: Int!) {
    sum(numbers: [$left, $right])
}
```

---

##### Using POST

> Standard form of interaction

```
POST /graphql
Content-Type: application/json
```

```json
{
    "query": "{\n    ping\n}",
}
```

_yields_

```json
{
    "data": {
        "ping": "pong"
    }
}
```

---

##### Using GET

> Cacheable through CDN and used in persisted queries

```
GET /graphql?query={ ping }
```

_yields_

```json
{
    "data": {
        "ping": "pong"
    }
}
```

---

##### What About `<Insert HTTP Method Here>`?

_No_, but...

---

# (￣～￣;) Where is GraphQL Used?

Everywhere!

- **Neo Financial**
- Netflix
- GitHub
- Dgraph
- Shopify
- Starbucks
- Lyft
- AirBnB

---

# (・・;)ゞ When Should You Use GraphQL?

Any time you are making an API!

- creating new APIs
- cleaning up legacy APIs
- messaging protocol(?)

---

# ヽ(ᵔ.ᵔ(・_・ )ゝ Why Should I Use GraphQL?

- strongly **(statically) typed**
- **no more DSLs** based on a case of NIHS
- **free validation** & less conditional boilerplate in code
- facilitates **communication across the full stack** (FE <-> BE)
- first class, **built-in documentation** (introspection)
- **no more "over-fetching"**, fetch only what you need
- federation **promotes microservices**

---

## Types, Scalars, and Enums; _Oh My_!

GraphQL offers a few core constructs for defining your schema types. The primary
one being `type`!

```graphql
type Character {
  name: String!
  appearsIn: [Episode!]!
  deceasedAt: String     # optional
}
```

---

### Scalars

The GraphQL type system is quite primitive, providing only a few constructs and
built-in scalar types of which there are five

- `Int`
- `Float`
- `String`
- `Boolean`
- `ID`

With most reference implementations offering the ability to define your own.

---

##### Custom Scalars

In the schema

```graphql
scalar Date
```

and in our implementation language of choice

```php
class Date extends Scalar
{
    public function serialize($value) {
        return $this->parseValue($value);
    }

    public function parseValue($value) {
        return $value->format('Y/M/D H:m:s');
    }

    public function parseLiteral(Node $node, ?array $variables = null) {
        return new DateTime($node->value);
    }
}
```

---

### Enums

Ultimately resolve as string values in the JSON output.

```graphql
enum Episode {
    NewHope
    Empire
    Jedi
    PhantomMenace
    CloneWars
    Sith
    ForceAwakens
    LastJedi
    Rise
}
```

---

### Interfaces

Ensure that types are defining a particular set of fields.

```graphql
interface Node {
    id: ID!
}

type Human implements Node {
  id: ID!
  name: String!
  friends: [Character!]!
  appearsIn: [Episode!]!
  starships: [Starship!]!
  totalCredits: Int
  homeworld: Planet!
}

type Droid implements Node {
  id: ID!
  name: String!
  friends: [Character!]!
  appearsIn: [Episode!]!
  primaryFunction: String
}
```

---

### Unions

The schema defines a union for search results

```graphql
union SearchResult = Human | Droid | Starship
```

our query might look something like this

```graphql
query Search {
    search(query: "skywalker") {
        ... on Human {
            name
            totalCredits
        }
        ... on Droid {
            name
            primaryFunction
        }
        ... on Starship {
            name
            owner {
                name
            }
        }
    }
}
```

---

### Inputs

What about complex inputs? Do `type`s work?

```graphql
input CustomerReview {
    stars: Int!
    commentary: String
}
```

And using it in a mutation would look like this with variables

```graphql
mutation ReviewEpisode($ep: Episode!, $review: CustomerReview!) {
  createReview(episode: $ep, review: $review) {
    stars
    commentary
  }
}
```

```json
{
  "ep": "JEDI",
  "review": {
    "stars": 5,
    "commentary": "This is a great movie!"
  }
}
```

---

## First-Class Documentation

No more Swagger, OpenAPI, et al.

```graphql
type User {
    """
    The user's unique ID in the form of an
    [RFC4122](https://datatracker.ietf.org/doc/html/rfc4122) compliant UUIDv4.
    """
    id: ID!
    """
    A unique handle (or "username") chosen by the user at signup. This value
    **SHOULD NOT** be relied upon as it can change at anytime at the descretion
    of the user.
    """
    alias: String!
    """
    A UNIX timestamp indicating when the user joined the service. See also
    `verifiedAt`.
    """
    joinedAt: Int!
}
```

---

# ┐('～`;)┌ Resolvers

> How do we wire data up to the schema?

- represents a node/field/edge on the graph
- resolvers are _basically_ (though not really) controllers
- (depending on reference implementation) they all receive the same 4, positional arguments

```php
function (mixed $parent, array $args, Context $context, ResolveInfo $info) {
    return null;
};
```

1. `parent`   – parent ("root") data for your current resolver
2. `args`     – all input arguments for your resolver field
3. `context`  – shared information between _all_ resolvers
4. `info`     – information about the query and current location in the graph

---

# ..・ヾ(。＞＜)シ Demo Time

### ‿︵‿︵‿︵‿ヽ(°□° )ノ︵‿︵‿︵‿︵

---

# ٩(× ×)۶ The n+1 Problem

Imagine your graph has the ability to query a list of comments and their author

```graphql
query Comments {
    comments {
        author {
            name
        }
    }
}
```

If we receive `n` comments back, the query executor will invoke `n` database
queries for the author, resulting in `n+1` calls. 

_dataloader to the rescue!_

```js
const keys = [1, 5, 3, 5, 5] // each resolver contributes a key to load
const load = (id) => {       // the loader will excute once and use the map to resolve fields
    authors.in(keys).reduce((authors, author) => {
        return {
            ...authors,
            [author.id]: author,
        }
    }, {})
}
```

---

# Too Many Requests

> Won't someone think of the ~~children~~ network!

##### Network Logs

| Name     | Status | Type | Size  | Time  |
|:---------|:-------|:-----|:------|:------|
| /graphql | 200    | POST |  869b | 199ms |
| /graphql | 200    | POST | 1.1kb | 229ms |
| /graphql | 200    | POST | 5.5kb | 211ms |
| /graphql | 200    | POST |  910b | 174ms |
| /graphql | 200    | POST | 2.2kb | 215ms |
| /graphql | 200    | POST | 1.1kb | 180ms |
| /graphql | 200    | POST |  778b | 293ms |
| /graphql | 200    | POST | 1.2kb | 301ms |
| /graphql | 200    | POST |  809b | 191ms |
| /graphql | 200    | POST | 6.0kb | 251ms |

Batch those queries!

```
POST /graphql
Content-Type: application/json
```

```json
[
    { "query": "query Profile($id: ID!) { ... }", "variables": { "id": "1" } },
    { "query": "query Followers($id: ID!) { ... }", "variables": { "id:" "1" } }
]
```

---

# (b ᵔ▽ᵔ)b Code Generation

- front-end (typescript)
- documentation, editor auto-complete

---

# ~(>_<~) Authentication & Authorization

In GraphQL, auth (authorization _and_ authentication) is a completely separate
layer and is meant to be implemented by you based on your business & technical
requirements;

```
Protocol:  |    HTTP     |
           ------ v ------
           |    Auth     |  <- B.Y.O.B
           ------ v ------
     Api:  |   GraphQL   |
           ------ v ------
           |     App     |
```

> _the same applies for for pagination and query constraints!_

---

# .･ﾟﾟ･(／ω＼)･ﾟﾟ･. Error Handling

- HTTP errors are different than GraphQL errors
    - e.g., 404 should only happen if your GraphQL route cannot be found
- graph will resolve as much as it can

---

# (*￣▽￣)b Microservices & Federation

- monolith has 1 giant schema (schema stitching)
- compose a graph from many microservices (gateway)
    - can load balance each service independently
    - gateway can be loadbalanced too!
- we use federation at Neo

`<load-balancing.png>`

`<fedartion.png>`

---

# The Hairy Parts

- learning curve – totally different paradigm
- performance – schema & query parsing overhead
- messy/unorganized – requires (cross-team) discipline
- familiarity for public APIs might be low

---

# ┐( ˘_˘ )┌ Recap/Key things

- there is no `SELECT *`, only fields
- resolvers are functions which return (scalar or structured) data for a field in the graph
- the schema represents your business domain, not the application (business language, not jargon)
- federation lets you compose a graph where each service is a contributer to the graph

---

# (⌐■_■) Thank You!

You can find the code for this talk on GitHub at

- [jordanbrauer/fullstack-mb-graphql](https://github.com/jordanbrauer/fullstack-mb-graphql)
  - _this slide deck can be found here ^ in_ `SLIDES.md`

and once again, you can find me on

- _GitHub_ as [@jordanbrauer](https://github.com/jordanbrauer)
- _Twitter_ as [@jorbandrauer](https://twitter.com/jorbandrauer)
- _Instagram_ as [@jorbandrauer](https://instagram.com/jorbandrauer)

---

# (・・ )? Questions

...

## Resources

- [Thinking in Graphs](https://graphql.org/learn/thinking-in-graphs/)
- [Introduction to GraphQL](https://graphql.org/learn/)
- [Global Object Identification](https://graphql.org/learn/global-object-identification/)

---

## Pop Quiz!

GraphQL allows you to query your database from your browser? (T/F)

---

## Pop Quiz!

REST is a terrible way to create an API (T/F)

---

## Pop Quiz!

What is the n+1 problem?

---

## Pop Quiz!

Apollo and GraphQL are the same (T/F)

---

_Thank You_
