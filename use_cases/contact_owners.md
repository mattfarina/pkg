# As a developer, I need to contact the maintainers of a package

There are cases where contacting a package owner directly, and not via
some publically facing issue queue (like GitHub or Launchpad) is
desirable.

Foremost may be the case where a severe security flaw needs to be
communicate to maintainers before it is publically disclosed.

## Notes
- The most reasonable and common communication medium is email.


## Solution
Package metadata includes an optional, but highly recommended, section
for declaring maintainers or persons of responsibility.

```json
{
    "name": "github.com/Masterminds/glide",
    "authors": [
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
