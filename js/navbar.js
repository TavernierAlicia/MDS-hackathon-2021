function createNav() {
  document.write(`
    <!DOCTYPE html>
    <html lang="fr">
      <head>
        <meta charset="utf-8">
        <link rel="icon" type="image/png" href="../pictures/favicon.png" />
      </head>
      <body>
        <nav onmouseleave=addClassNav() onmouseenter=removeClassNav()>
          <a class="index" href="/"><img src="/pictures/logo.png" alt="logo"></a>
          <div class="navbar-right-items">
	          <button id="menuButton">Menu</button>
            <ul>
              <li class="nav-item">
                <form method="POST" name="search" action="/" method="post" id="form" class="nav-search">
                  <input autocomplete="off" type="search" class="searchBar" id="site-search" name="keyword" placeholder="ex: Comment obtenir les APLs?">
                  <input type="submit" id="sendform" value="Chercher" hidden>
                </form>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/housing">Logements</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/rights">Mes droits</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/faq">FAQ</a>
              </li>
            </ul>
          </div>
        </nav>
      </body>
    </html>
  `);
}

createNav();