# As a developer, I need to contact the maintainers of a package

There are cases where contacting a package owner directly, and not via
some publicly facing issue queue (like GitHub or Launchpad) is
desirable.

Foremost may be the case where a severe security flaw needs to be
communicate to maintainers before it is publicly disclosed.

## Notes
- The most reasonable and common communication medium is email.
- An owner could be a person, a collection of people, a organization
  (e.g., a company). The contact information should be able to reflect
  any of these.

## Solution
Package metadata includes an optional, but highly recommended, section
for declaring maintainers or persons of responsibility.

```json
{
    "name": "github.com/Masterminds/glide",
    "owners": [
        {
            "name": "Matt Butcher",
            "email": "something@example.com",
            "homepage": "http://technosophos.com"
        },
        {
            "name": "Matt Farina",
            "email": "other@example.com",
            "homepage": "http://mattfarina.com"
        }
    ],
    "imports": []
}
```
