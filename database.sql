-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Feb 19, 2021 at 01:17 AM
-- Server version: 5.7.26
-- PHP Version: 7.2.18

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `studenth`
--

-- --------------------------------------------------------

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(250) NOT NULL,
  `main_picture` varchar(250) DEFAULT NULL,
  `creation_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `article`
--

INSERT INTO `article` (`id`, `title`, `main_picture`, `creation_date`) VALUES
(1, 'La garantie Visale.', 'https://www.actionlogement.fr/sites/als/files/images/Actualites/2018/vignette_actu_visale.jpg', '2021-02-10 23:00:00'),
(2, 'La prime d\'activité', 'https://www.vosgelis.fr/wp-content/uploads/2019/08/Prime-activit%C3%A9.jpg', '2021-02-16 23:00:00'),
(3, 'Les allocations familliales', 'https://www.maisondeservicesaupublic.fr/sites/default/files/public/thumbnails/image/caf.png', '2021-02-10 23:00:00'),
(4, 'Où se faire soigner?', 'https://img.aws.la-croix.com/2017/05/01/1200843666/Medecin-stethoscope_0_729_486.jpg', '2021-02-16 23:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `article_content`
--

DROP TABLE IF EXISTS `article_content`;
CREATE TABLE IF NOT EXISTS `article_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `low_title` varchar(250) DEFAULT NULL,
  `text` text,
  `picture` varchar(250) DEFAULT NULL,
  `article_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `article_content`
--

INSERT INTO `article_content` (`id`, `low_title`, `text`, `picture`, `article_id`) VALUES
(1, 'Qui est bénéficiaire?', 'Cras at ipsum egestas, pretium dolor ut, dictum elit. Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed. Donec sed condimentum sapien. Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis. Interdum et malesuada fames ac ante ipsum primis in faucibus. In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, non tincidunt purus consequat eu.', 'https://www.visale.fr/wp-content/uploads/2019/09/illu_homepage_200000.png', 1),
(2, 'De quels documents ai-je besoin?', 'Cras at ipsum egestas, pretium dolor ut, dictum elit. Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed. Donec sed condimentum sapien. Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis. Interdum et malesuada fames ac ante ipsum primis in faucibus. In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, \\n - non tincidunt purus consequat eu.\r\ntortor massa, ut aliquam enim commodo non. \\n \r\n- Vivamus pulvinar, est nec suscipit efficitur, mauris tortor tempor sem, sit amet maximus nisi purus a augue. \\n \r\n- Donec turpis lorem, ornare at euismod quis, varius eget dolor. Donec sit amet tellus posuere, efficitur enim vitae, consequat sapien. Nullam egestas elit ac fermentum mollis. \\n\r\n- Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. \r\n\r\nProin et vestibulum es', NULL, 1),
(3, 'Les démarches à faire', 'Phasellus molestie sem at turpis facilisis, et sagittis libero varius. Integer bibendum nunc sed ipsum ultricies finibus id eu nisl. Duis eu justo pretium, semper felis ut, laoreet nunc. Quisque nec ex et nulla condimentum condimentum ut a turpis. Donec erat turpis, fringilla et erat a, fermentum laoreet purus. Donec egestas lacus eget ultricies pulvinar. Aenean viverra fermentum ipsum ut scelerisque. Sed vitae bibendum nibh. Phasellus erat augue, commodo et purus nec, elementum aliquam mi.', 'https://www.caf.fr/sites/default/files/caf/428/Images/visale/visale%20image%20article%20VDF.jpg', 1),
(4, 'À qui s\'adresse la prime d\'activité?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur condimentum leo volutpat ultrices consectetur. Vestibulum aliquam commodo lacus. Donec mattis faucibus egestas. Mauris condimentum est ac purus vestibulum, et dignissim ipsum dictum. Maecenas ultricies metus diam, eget pretium eros suscipit sit amet. Ut lorem elit, malesuada et rhoncus tincidunt, eleifend quis nisl. Cras dapibus at nisi vel auctor.\r\n\\n\r\nVivamus tellus ligula, consequat id urna at, consequat fringilla mauris. Nulla odio odio, viverra faucibus massa ut, consequat aliquet lorem. Pellentesque non leo lacinia, laoreet risus eget, vestibulum lorem. Donec tincidunt lacus augue, at luctus ex facilisis luctus. Aenean tempor tortor massa, ut aliquam enim commodo non. Vivamus pulvinar, est nec suscipit efficitur, mauris tortor tempor sem, sit amet maximus nisi purus a augue. Donec turpis lorem, ornare at euismod quis, varius eget dolor. Donec sit amet tellus posuere, efficitur enim vitae, consequat sapien. Nullam egestas elit ac fermentum mollis. Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. Proin et vestibulum est.\r\n\\n\r\nPhasellus molestie sem at turpis facilisis, et sagittis libero varius. Integer bibendum nunc sed ipsum ultricies finibus id eu nisl. Duis eu justo pretium, semper felis ut, laoreet nunc. Quisque nec ex et nulla condimentum condimentum ut a turpis. Donec erat turpis, fringilla et erat a, fermentum laoreet purus. Donec egestas lacus eget ultricies pulvinar. Aenean viverra fermentum ipsum ut scelerisque. Sed vitae bibendum nibh. Phasellus erat augue, commodo et purus nec, elementum aliquam mi.', NULL, 2),
(5, 'Comment l\'obtenir?', 'Cras at ipsum egestas, pretium dolor ut, dictum elit. Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed. Donec sed condimentum sapien. Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis. Interdum et malesuada fames ac ante ipsum primis in faucibus. In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, non tincidunt purus consequat eu.', 'https://www.toutsurmesfinances.com/impots/wp-content/uploads/sites/6/2015/12/simulateur-prime-dactivite.jpg', 2),
(6, 'Les CHU', 'Cras at ipsum egestas, pretium dolor ut, dictum elit. Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed. Donec sed condimentum sapien. Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis. Interdum et malesuada fames ac ante ipsum primis in faucibus. In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, \\n - non tincidunt purus consequat eu.\r\ntortor massa, ut aliquam enim commodo non. \\n \r\n- Vivamus pulvinar, est nec suscipit efficitur, mauris tortor tempor sem, sit amet maximus nisi purus a augue. \\n \r\n- Donec turpis lorem, ornare at euismod quis, varius eget dolor. Donec sit amet tellus posuere, efficitur enim vitae, consequat sapien. Nullam egestas elit ac fermentum mollis. \\n\r\n- Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. \r\n\r\nProin et vestibulum es', 'https://media.lesechos.com/api/v1/images/view/5bd1d3158fe56f2040666482/1280x720/2147295-soigner-lhopital-et-ses-finances-178356-1.jpg', 4),
(7, 'SOS Médecins', 'Phasellus molestie sem at turpis facilisis, et sagittis libero varius. Integer bibendum nunc sed ipsum ultricies finibus id eu nisl. Duis eu justo pretium, semper felis ut, laoreet nunc. Quisque nec ex et nulla condimentum condimentum ut a turpis. Donec erat turpis, fringilla et erat a, fermentum laoreet purus. Donec egestas lacus eget ultricies pulvinar. Aenean viverra fermentum ipsum ut scelerisque. Sed vitae bibendum nibh. Phasellus erat augue, commodo et purus nec, elementum aliquam mi.', '', 4),
(8, 'Quels services?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur condimentum leo volutpat ultrices consectetur. Vestibulum aliquam commodo lacus. Donec mattis faucibus egestas. Mauris condimentum est ac purus vestibulum, et dignissim ipsum dictum. Maecenas ultricies metus diam, eget pretium eros suscipit sit amet. Ut lorem elit, malesuada et rhoncus tincidunt, eleifend quis nisl. Cras dapibus at nisi vel auctor.\r\n\\n\r\nVivamus tellus ligula, consequat id urna at, consequat fringilla mauris. Nulla odio odio, viverra faucibus massa ut, consequat aliquet lorem. Pellentesque non leo lacinia, laoreet risus eget, vestibulum lorem. Donec tincidunt lacus augue, at luctus ex facilisis luctus. Aenean tempor tortor massa, ut aliquam enim commodo non. Vivamus pulvinar, est nec suscipit efficitur, mauris tortor tempor sem, sit amet maximus nisi purus a augue. Donec turpis lorem, ornare at euismod quis, varius eget dolor. Donec sit amet tellus posuere, efficitur enim vitae, consequat sapien. Nullam egestas elit ac fermentum mollis. Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. Proin et vestibulum est.\r\n\\n\r\nPhasellus molestie sem at turpis facilisis, et sagittis libero varius. Integer bibendum nunc sed ipsum ultricies finibus id eu nisl. Duis eu justo pretium, semper felis ut, laoreet nunc. Quisque nec ex et nulla condimentum condimentum ut a turpis. Donec erat turpis, fringilla et erat a, fermentum laoreet purus. Donec egestas lacus eget ultricies pulvinar. Aenean viverra fermentum ipsum ut scelerisque. Sed vitae bibendum nibh. Phasellus erat augue, commodo et purus nec, elementum aliquam mi.', 'https://stickeramoi.com/17208-large_default/sticker-tirelire-cochon-rose-argent.jpg', 3),
(9, 'Comment postuler?', 'Cras at ipsum egestas, pretium dolor ut, dictum elit. Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed. Donec sed condimentum sapien. Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis. Interdum et malesuada fames ac ante ipsum primis in faucibus. In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, non tincidunt purus consequat eu.', 'https://www.toutsurmesfinances.com/impots/wp-content/uploads/sites/6/2015/12/simulateur-prime-dactivite.jpg', 3);

-- --------------------------------------------------------

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
CREATE TABLE IF NOT EXISTS `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pseudo` varchar(100) NOT NULL DEFAULT 'Anonymous',
  `content` text NOT NULL,
  `creation_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `article_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `comment`
--

INSERT INTO `comment` (`id`, `pseudo`, `content`, `creation_date`, `article_id`) VALUES
(1, 'Germaine.H', 'Super intéressant!', '2021-02-18 23:38:28', 1),
(6, 'Ginette.L', 'Cet article m\'a beaucoup aidée!', '2021-02-19 01:14:33', 2),
(7, 'Estelle.P', 'Super article!', '2021-02-19 01:14:43', 2);

-- --------------------------------------------------------

--
-- Table structure for table `faq`
--

DROP TABLE IF EXISTS `faq`;
CREATE TABLE IF NOT EXISTS `faq` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question` text NOT NULL,
  `answer` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `faq`
--

INSERT INTO `faq` (`id`, `question`, `answer`) VALUES
(1, 'Mes enfants de 19 et 21 ans habitent ensemble, ils sont colocataires. Comment faire la\r\ndemande d\'aide au logement ?', 'Chacun de vos enfants doit faire sa propre demande d\'aide au logement en tant que colocataire\r\nsur caf.fr. Chacun aura son dossier, son numéro allocataire personnel, ainsi que son mot de passe\r\npersonnel.'),
(2, 'Vivamus tellus ligula, consequat id urna at, consequat fringilla mauris?', 'Nulla odio odio, viverra faucibus massa ut, consequat aliquet lorem. Pellentesque non leo lacinia, laoreet risus eget, vestibulum lorem. Donec tincidunt lacus augue, at luctus ex facilisis luctus. Aenean tempor tortor massa, ut aliquam enim commodo non. Vivamus pulvinar, est nec suscipit efficitur, mauris tortor tempor sem, sit amet maximus nisi purus a augue. Donec turpis lorem, ornare at euismod quis, varius eget dolor. Donec sit amet tellus posuere, efficitur enim vitae, consequat sapien. Nullam egestas elit ac fermentum mollis. Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. Proin et vestibulum est.'),
(3, 'Lorem ipsum dolor sit amet, consectetur adipiscing elit?', 'Curabitur condimentum leo volutpat ultrices consectetur. Vestibulum aliquam commodo lacus. Donec mattis faucibus egestas. Mauris condimentum est ac purus vestibulum, et dignissim ipsum dictum. Maecenas ultricies metus diam, eget pretium eros suscipit sit amet. Ut lorem elit, malesuada et rhoncus tincidunt, eleifend quis nisl. Cras dapibus at nisi vel auctor.'),
(4, 'Nullam egestas elit ac fermentum mollis. Curabitur commodo gravida nisi, et tincidunt lectus efficitur sit amet. Praesent vel odio felis. Proin et vestibulum est?', 'Phasellus molestie sem at turpis facilisis, et sagittis libero varius. Integer bibendum nunc sed ipsum ultricies finibus id eu nisl. Duis eu justo pretium, semper felis ut, laoreet nunc. Quisque nec ex et nulla condimentum condimentum ut a turpis. Donec erat turpis, fringilla et erat a, fermentum laoreet purus. Donec egestas lacus eget ultricies pulvinar. Aenean viverra fermentum ipsum ut scelerisque. Sed vitae bibendum nibh. Phasellus erat augue, commodo et purus nec, elementum aliquam mi.'),
(5, 'Cras at ipsum egestas, pretium dolor ut, dictum elit?', 'Mauris sed scelerisque odio. Nullam condimentum, ante vitae dapibus consectetur, mauris justo imperdiet lorem, in laoreet nibh quam vitae nisi. Mauris facilisis tellus vitae felis malesuada, quis aliquam sapien rutrum. Vestibulum feugiat massa sit amet porttitor ultricies. Maecenas finibus massa tortor, ac cursus nisl vulputate sed.'),
(6, 'Donec sed condimentum sapien?', 'Donec sed tellus est. Ut ullamcorper lectus vitae ex gravida, id vehicula leo lacinia. Aenean volutpat metus at nulla ultrices, in aliquet lacus congue. Etiam libero nisl, suscipit ac porttitor vel, finibus sit amet lorem. Cras non urna in ante convallis venenatis.'),
(7, 'Interdum et malesuada fames ac ante ipsum primis in faucibus?', ' In eros magna, rhoncus sit amet elit vehicula, dignissim lobortis lacus. Nunc vel odio a tellus porttitor lobortis. Aenean lacinia pellentesque purus, non tincidunt purus consequat eu.'),
(8, 'Donec nunc felis, dapibus quis posuere ut, vestibulum eu libero?', 'Morbi et turpis erat. Proin malesuada, justo sed elementum ultrices, orci ligula pulvinar massa, eget suscipit ex libero vel leo. Nam neque nibh, dictum sed venenatis a, efficitur lobortis lectus. Proin in purus at sem maximus placerat. Praesent volutpat ex non consequat fermentum. Cras non sollicitudin tortor, vel venenatis enim. Donec turpis metus, varius ut elit sit amet, hendrerit suscipit urna. Maecenas ac leo quis leo tincidunt egestas. In posuere metus lectus, eu ultrices purus vestibulum sit amet.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `article`
--
ALTER TABLE `article` ADD FULLTEXT KEY `title` (`title`);

--
-- Indexes for table `article_content`
--
ALTER TABLE `article_content` ADD FULLTEXT KEY `text` (`text`);
ALTER TABLE `article_content` ADD FULLTEXT KEY `low_title` (`low_title`);

--
-- Indexes for table `faq`
--
ALTER TABLE `faq` ADD FULLTEXT KEY `question` (`question`);
ALTER TABLE `faq` ADD FULLTEXT KEY `answer` (`answer`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
