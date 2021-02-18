function createFooter() {
  document.write(`
  <!DOCTYPE html>
  <html lang="fr">
    <head>
      <meta charset="utf-8">
      <title>Studenth - Accueil</title>
    </head>
    <body>
    </body>
    <footer>
      <div class="container">
      	<div class="wrapper">
        <div class="footer-left">
          <div class="footer-line">
            <a href="/faq">FAQ</a>
          </div>
        </div>
        <div class="footer-middle">
          <div class="footer-line">
            <a href="/legal"> Mentions légales </a>
          </div>
        </div>

        <div class="footer-right">
          <div class="footer-line">
            <a href="#"><img class=social src="../pictures/fb.png" alt="logo de facebook"></a>
            <a href="#"><img class=social src="../pictures/insta.png" alt="logo d'instagram"></a>
            <a href="#"><img class=social src="../pictures/twitter.png" alt="logo de twitter"></a>
          </div>
        </div>
      </div>
      </div>
      <div class="cp">
        <p>Studenth - Copyright 2021</p>
      </div>
    </footer>
  </html>
  `);

}

createFooter();
