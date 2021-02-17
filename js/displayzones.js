var zones = [
  // [value, display]
  ["tous", "tous"],
  ["Aix-en-Provence", "aix-en-provence"],
  ["Albi", "albi"],
  ["AMIENS CAMPUS", "amiens campus"],
  ["AMIENS CENTRE", "amiens centre"],
  ["ANGERS", "angers"],
  ["Angoulême", "angoulême"],
  ["Annecy - Domaine Universitaire d'Annecy-Le-Vieux", "annecy - domaine universitaire d'annecy-le-vieux"],
  ["Arras", "arras"],
  ["Avignon", "avignon"],
  ["Belfort", "belfort"],
  ["Besançon - Centre ville", "besançon - centre ville"],
  ["Blois", "blois"],
  ["Bobigny", "bobigny"],
  ["Bordeaux", "bordeaux"],
  ["Boulogne sur Mer", "boulogne sur mer"],
  ["Bourges", "bourges"],
  ["Brest", "brest"],
  ["Caen", "caen"],
  ["Campus Charles Mérieux - Gerland", "campus charles mérieux - gerland"],
  ["Campus Charles Mérieux - Manufacture des Tabacs", "campus charles mérieux - manufacture des tabacs"],
  ["Campus Charles Mérieux - Site des Quais / Manufacture des Tabacs", "campus charles mérieux - site des quais / manufacture des tabacs"],
  ["Campus Grimaldi", "campus grimaldi"],
  ["Campus Lyon Santé Est / Campus Porte des Alpes Bron", "campus lyon santé est / campus porte des alpes bron"],
  ["Campus Lyon Tech La Doua", "campus lyon tech la doua"],
  ["Chambéry - Campus Universitaire de Jacob", "chambéry - campus universitaire de jacob"],
  ["CHARLEVILLE - Centre", "charleville - centre"],
  ["Chartres", "chartres"],
  ["Châteauroux et Issoudun", "châteauroux et issoudun"],
  ["Clermont-Ferrand", "clermont-ferrand"],
  ["COMPIEGNE", "compiegne"],
  ["Creteil", "creteil"],
  ["Côte Basque", "côte basque"],
  ["DIJON CAMPUS", "dijon campus"],
  ["Dordogne", "dordogne"],
  ["Epinal", "epinal"],
  ["ESSONNE", "essonne"],
  ["Grenoble - Dans l'agglomération", "grenoble - dans l'agglomération"],
  ["Grenoble - Domaine Universitaire", "grenoble - domaine universitaire"],
  ["Grenoble - En ville", "grenoble - en ville"],
  ["Guadeloupe", "guadeloupe"],
  ["HAUTS DE SEINE", "hauts de seine"],
  ["Herouville Saint Clair", "herouville saint clair"],
  ["La Garde", "la garde"],
  ["La Rochelle", "la rochelle"],
  ["Landes", "landes"],
  ["Le Bourget du Lac - En ville", "le bourget du lac - en ville"],
  ["Le Havre", "le havre"],
  ["Les Cezeaux", "les cezeaux"],
  ["Lille (Moulins)", "lille (moulins)"],
  ["Limoges (87)", "limoges"],
  ["Lyon 5ème", "lyon 5ème"],
  ["Marne la Vallee", "marne la vallee"],
  ["Marseille", "marseille"],
  ["Martinique", "martinique"],
  ["Metz", "metz"],
  ["Mont Saint Aignan", "mont saint aignan"],
  ["Montpellier", "montpellier"],
  ["Mulhouse", "mulhouse"],
  ["Nancy", "nancy"],
  ["NANTES", "nantes"],
  ["Nice", "nice"],
  ["Nord", "nord"],
  ["Nîmes", "nîmes"],
  ["Orléans", "orléans"],
  ["Paris 05", "paris 05"],
  ["Paris 11", "paris 11"],
  ["Paris 12", "paris 12"],
  ["Paris 13", "paris 13"],
  ["Paris 14", "paris 14"],
  ["Paris 15", "paris 15"],
  ["Paris 18", "paris 18"],
  ["Paris 19", "paris 19"],
  ["Paris 20", "paris 20"],
  ["Pau", "pau"],
  ["Perpignan", "perpignan"],
  ["Pessac", "pessac"],
  ["Poitiers Campus", "poitiers campus"],
  ["Poitiers Centre Ville", "poitiers centre ville"],
  ["REIMS - Campus Lettres", "reims - campus lettres"],
  ["REIMS - Campus Sciences", "reims - campus sciences"],
  ["Rennes Centre", "rennes centre"],
  ["Rennes est-Beaulieu", "rennes est-beaulieu"],
  ["Rennes ouest-Villejean", "rennes ouest-villejean"],
  ["Roubaix", "roubaix"],
  ["Rouen", "rouen"],
  ["SAINT-ETIENNE", "saint-etienne"],
  ["St Denis", "st denis"],
  ["Strasbourg", "strasbourg"],
  ["Sénart", "sénart"],
  ["Talence", "talence"],
  ["Toulouse Centre", "toulouse centre"],
  ["Toulouse sud-est", "toulouse sud-est"],
  ["Tours", "tours"],
  ["VAL D'OISE", "val d'oise"],
  ["Valence - En ville", "valence - en ville"],
  ["Valenciennes (Domaine Universitaire)", "valenciennes (domaine universitaire)"],
  ["Vandoeuvre-lès-Nancy", "vandoeuvre-lès-nancy"],
  ["Villeneuve d’Ascq (Cité scientifique)", "villeneuve d’ascq (cité scientifique)"],
  ["Villeneuve d’Ascq (Pont de bois)", "villeneuve d’ascq (pont de bois)"],
  ["YVELINES", "yvelines"],
]




function displayZones() {

  for(var i = 0; i < zones.length; i++) {
    var zone = zones[i];
    var newOption = document.createElement("option")

    // insert in options menu
    if (zone[0] == "tous") {
      newOption.setAttribute("selected", "selected")
    }
    newOption.setAttribute("value", zone[0])
    newOption.innerHTML = zone[1]

    // add to select field
    document.getElementsByTagName("select")[0].appendChild(newOption)
  }
}

displayZones()



















