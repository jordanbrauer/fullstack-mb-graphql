---
author: Jordan Brauer
date: MMMM dd, YYYY
---

# (^０^)ノHello!

My name is Jordan Brauer

I've been doing software for a few years —

- ~10 years unproffesionally as a young lad
- ~6 years professionally


1. 2.0 Ma & pa, freelance
2. 3.5 Pepper (Tipping Canoe)
3. 0.5 Gubagoo
4. 0.5 Neo Financial

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

## ┐( ˘_˘ )┌

- strongly (statically) typed
- no more DSLs based on a case of NIHS
- free validation & less conditional boilerplate in code
- facilitates communication across the full stack (FE <-> BE)
- introspection
- first class documentation
- fetch only what you need ("over-fetching")
- federation promotes microservices

---

## Types, Scalars, and Enums; _Oh My_!

```graphql
type Character {
  name: String!
  appearsIn: [Episode!]!
}
```

---

### Scalars

**Five** built-in scalar types

- `Int`
- `Float`
- `String`
- `Boolean`
- `ID`

With most reference implementations offering the ability to define your own

##### Custom Scalars

```graphql
scalar Date
```

```php
class Date extends Scalar
{
    public function serialize($value)
    {
        return $this->parseValue($value);
    }

    public function parseValue($value)
    {
        return $value;
    }

    public function parseLiteral(Node $node, ?array $variables = null)
    {
        return $node->value;
    }
}
```

---

### Enums

```graphql
enum Episode {
    NEWHOPE
    EMPIRE
    JEDI
}
```

---

### Interfaces

```graphql
interface Character {
    id: ID!
    name: String!
    friends: [Character]!
    appearsIn: [Episode]!
}

type Human implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  starships: [Starship]
  totalCredits: Int
}

type Droid implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  primaryFunction: String
}
```

---

### Unions

```graphql
union SearchResult = Human | Droid | Starship
```

---

### Inputs

```graphql
input CustomerReview {
    stars: Int!
    commentary: String
}
```

client uses it like

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

No more Swagger!

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

    # ...
}
```

---

# ┐('～`;)┌ Resolvers

- resolvers are basically (though not really) controllers
- they all receive the same 4, positional arguments
- represents a node/field/edge on the graph

```php
function (mixed $parent, array $args, Context $context, ResolveInfo $info) {
    return null;
};
```

- parent (root)
- arguments
- context
- info

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

---

# (b ᵔ▽ᵔ)b Code Generation

- front-end (typescript)
- back-end (Go)
- documentation, editor auto-complete

---

# .･ﾟﾟ･(／ω＼)･ﾟﾟ･. Error Handling

- HTTP errors are different than GraphQL errors
    - e.g., 404 should only happen if your GraphQL route cannot be found

---

# ٩(× ×)۶ The n+1 Problem

- dataloader pattern (made by Facebook)
- essentially a buffer of resource IDs to then be loaded all at once

---

# Batched Queries

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

# (*￣▽￣)b Microservices & Federation

- monolith has 1 giant schema (schema stitching)
- compose a graph from many microservices (gateway)

---

# The Hairy Parts

- learning curve – totally different paradigm
- performance – schema & query parsing overhead
- messy/unorganized – requires (cross-team) discipline

---

# ..・ヾ(。＞＜)シ Demo Time

### ‿︵‿︵‿︵‿ヽ(°□° )ノ︵‿︵‿︵‿︵

---

# Recap/Key things

- there is no `SELECT *`, only fields
- resolvers are functions which return (scalar or structured) data for a field in the graph
- the schema represents your business domain, not the application (business language, not jargon)
- federation & microservices let you compose a graph where each service is a contributer to the graph

---

# (⌐■_■) Thank You!

You can find the code for this talk on GitHub at

- [craig/rest](https://github.com)
- [jordanbrauer/fullstack-mb-graphql](https://github.com/jordanbrauer/fullstack-mb-graphql)
  - _this slide deck can be found here ^ in_ `SLIDES.md`
- [jordanbrauer/fullstack-mb-ui](https://github.com/jordanbrauer/fullstack-mb-ui)

and once again, I am on –

- _GitHub_ as [@jordanbrauer](https://github.com/jordanbrauer)
- _Twitter_ as [@jorbandrauer](https://twitter.com/jorbandrauer)
- _Instagram_ as [@jorbandrauer](https://instagram.com/jorbandrauer)

## Resources

- [Thinking in Graphs](https://graphql.org/learn/thinking-in-graphs/)
- [Introduction to GraphQL](https://graphql.org/learn/)
- [Global Object Identification](https://graphql.org/learn/global-object-identification/)

## (・・ )? Questions

---

> REMOVE ME BEFORE FINAL DRAFT

# Speaker Notes

- show a hello world example and explain resolvers;
- show transition from REST to GraphQL by masking the REST API with GraphQL;
- monolith vs. microservice & service-to-services calls/direct DB;

- it's a spec to be implemented in whatever language you choose;
- what is this solving?;
  - vs. REST approach (over fetch, fetching many resources);
  - scaling the API (microservices & service to service calls, etc.);
  - start REST and show what problems GraphQL solves;
  - services all speak common GraphQL

## Advanced Topics

- apollo adds to GraphQL with plugins and federation;
- subscriptions (realtime w/ socket);
