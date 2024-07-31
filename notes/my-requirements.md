# api
- server side rendering [ x ]
- library-free http server framework [ x ]
- JWT middleware
- secret management


# email
- email users for delivery/shipping
- client to external service or possibly use free (poc gmail's?)
- cron here or separate?


# crawler
- accept url, sanitize, validate, parse html [ x ]
- common template attributes (manually validated)
- rate limited? not public?
- start with target.com [ x ]
- MVP is this is self-generated instead


# web
- typescript web client?
- or simple http
- styling
- MVP potentially just server-side rendering [ x ]


# database
- sqlite simple file mvp
- postgres?
- url or content storage?
- user accounts

# worker communications
- message broker?
- k8s service direct? grpc bonus? (zero durability)
- dead simple message distribution maybe...


# accounts
- consider authentication/authorization through managed service? auth0?
- or internal IdP with JWT


## Stretch Goals or Future Improvement
- rate limiting maybe in service/code (but can be done at http level though (e.g., nginx, cloudflare))
- categories for cached items to make the site more like a real ecommerce thing