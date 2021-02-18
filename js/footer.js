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


        <!-- LEFT -->
        <div class="footer-left">
          <div class="footer-line">
            <a class="logofoot" href="/"><img src="../pictures/logo.png"></a>
          </div>
        </div>


        <!-- MIDDLE -->
        <div class="footer-middle">
          <div class="footer-line">
            <a href="#">Nous Contacter</a>
          </div>
          <div class="footer-line">
            <a href="#" class="lighttitle" ><img class="picto" src="../pictures/tel.png"> 06 47 56 78 34</a>
          </div>
          <div class="footer-line">
            <a href="#" class="lighttitle"><img class="picto" src="../pictures/mail.png"> contact@studenth.fr</a>
          </div>
          <div class="footer-line">
            <p>Horaires</p>
            <p class="lighttitle">9h - 12h30 / 13h30 - 18h</p>
            <p class="lighttitle">Du lundi au vendredi sauf jours fériés</p>
          </div>
        </div>

        <!-- RIGHT -->
        <div class="footer-right">
          <div class="footer-line">
            <a href="/faq"><img class="picto" src="../pictures/arrow.png">FAQ</a>
          </div>
          <div class="footer-line">
            <a href="/legal"><img class="picto" src="../pictures/arrow.png">Mentions légales </a>
          </div>
          <div class="footer-line">
            <a href="#"><img class="picto" src="../pictures/arrow.png">Télémaintenance </a>
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
