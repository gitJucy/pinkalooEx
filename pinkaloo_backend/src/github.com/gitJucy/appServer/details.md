# Coding Exercise - Application Metadata API Server

## Requirements

- Build a RESTful API server for persisting application metadata
- An endpoint to persist application metadata (In memory is fine). The API must support JSON as a valid payload format.
- An endpoint to search application metadata and retrieve a list that matches the query parameters.
- Include tests if you feel it’s appropriate.

JSON payloads are provided. Two that should persist, and two that should error due to missing fields.

## Rules

Any software or open source library is fair game to help you solve this problem. The response from the server as well as the structure of the query endpoint is intentionally vague to allow for flexibility and innovation in your solution.

## Advice

This exercise is an opportunity for you to show off your passion and artisanship of your solution. Optimize for quality and reliability. If you feel your solution is missing a cool feature and you have time, have fun and add it. Make the solution your own, and show off your skills.

## What about the database?

You can if you want, but it's not necessary. Integrating with a database driver or ORM gives you less room to shine, and us less ability to evaluate your work.

## Example payloads

All fields in the payload are required. For illustration purposes, we have a few example payloads. One example payload where the maintainer email is not a valid email and another where the version is missing that should fail on submit and two that should be valid.

### Invalid Payloads

```json
{
  "title": "App w/ Invalid maintainer email",
  "version": "1.0.1",
  "maintainers": [
    {
      "name": "Firstname Lastname",
      "email": "apptwohotmail.com"
    }
  ],
  "company": "Pinkaloo Technologies Inc.",
  "website": "https://pinkaloo.com",
  "source": "https://github.com/pinkaloo/repo",
  "license": "Apache-2.0",
  "description": "### blob of markdown"
}
```

```json
{
  "title": "App w/ missing version",
  "maintainers": [
    {
      "name": "first last",
      "email": "email@hotmail.com"
    },
    {
      "name": "first last",
      "email": "email@gmail.com"
    }
  ],
  "company": "Company Inc.",
  "website": "https://website.com",
  "source": "https://github.com/company/repo",
  "license": "Apache-2.0",
  "description": "### blob of markdown"
}
```

### Valid Payloads

```json
{
  "title": "Valid App 1",
  "version": "0.0.1",
  "maintainers": [
    {
      "name": "firstmaintainer app1",
      "email": "firstmaintainer@hotmail.com"
    },
    {
      "name": "secondmaintainer app1",
      "email": "secondmaintainer@gmail.com"
    }
  ],
  "company": "Random Inc.",
  "website": "https://website.com",
  "source": "https://github.com/random/repo",
  "license": "Apache-2.0",
  "description": "### Interesting Title"
}
```

```json
{
  "title": "Valid App 2",
  "version": "1.0.1",
  "maintainers": [
    {
      "name": "AppTwo Maintainer",
      "email": "apptwo@hotmail.com"
    }
  ],
  "company": "Pinkaloo Technologies",
  "website": "https://pinkaloo.com",
  "source": "https://github.com/pinkaloo/repo",
  "license": "Apache-2.0",
  "description": "### Why app 2 is the best"
}
```
