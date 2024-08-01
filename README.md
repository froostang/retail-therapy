# retail-therapy


In the case of Retail Therapy, the idea could otherwise be described as an outlet for pretending to purchase items. The goal of creating fake purchases is to try and replicate the endorphin rush that occurs when you press the “checkout” button on a fun new item. The anticipation and excitement of awaiting your shiny new toy can often be almost as enjoyable as the item itself. Thus, Retail Therapy would be an outlet to attempt to trigger these reactions without the financial drawbacks.





## Usage

- If you want to run the app locally, be sure to have Go 1.21+ installed and `cd retail-therapy/api` then `make build` and `./build/api-service`.
- Navigate to `localhost:8080/shop` to see the site.
- For dev work: just run the command `air` while in /api to read from the `.air.toml`
  
Use the `/add-product` page to include links to add to your fake store. Then click the shop page after to see them populate. You can try adding them to cart, then click the checkout icon. If you enter fake payment information on the checkout page and hit complete purchase you should get some fun confetti and the cart will reset.
Important Note: Currently the `/add-product` functionality works only with specifice/sanitized store links (by design, scraping is hard and dangerous). Not 100% will scrape correctly.
