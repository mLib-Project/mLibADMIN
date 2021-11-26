package main

var books = []book{

	// ########################################################################
	// ########################################################################
	// ########################################################################
	// ########################################################################

	// Matematyka
	{category: 0, subcategory: "", name: "Matematyka", author: "AGH Kraków", source: "https://epodreczniki.open.agh.edu.pl/handbook/list?categories=1"},
	{category: 0, subcategory: "", name: "Matematyka", author: "KUL Lublin", source: "http://paruchl.webd.pl/mat/matematyka.pdf"},
	{category: 0, subcategory: "Wstęp do Matematyki", name: "Wstęp do matematyki", author: "P. Jędrzejewicz", source: "https://www-users.mat.umk.pl/~pjedrzej/wstep/wyklad.pdf"},
	{category: 0, subcategory: "Wstęp do Matematyki", name: "Wstęp do matematyki", author: "J. Kraszewski", source: "https://docer.pl/doc/evevx"},
	{category: 0, subcategory: "Wstęp do Matematyki", name: "Wstęp do matematyki współczesnej", author: "H. Rasiowa", source: "https://docer.pl/doc/xs00"},
	{category: 0, subcategory: "Algebra", name: "Algebra liniowa I - Twierdzenia", author: "T. Jurlewicz, Z. Skoczylas", source: "https://docer.pl/doc/nsxvssv"},
	{category: 0, subcategory: "Algebra", name: "Algebra liniowa I - Zadania", author: "T. Jurlewicz, Z. Skoczylas", source: "https://docer.pl/doc/n0x0n1s"},
	{category: 0, subcategory: "Algebra", name: "Algebra liniowa II - Twierdzenia", author: "T. Jurlewicz, Z. Skoczylas", source: "https://docer.pl/doc/esvse"},
	{category: 0, subcategory: "Algebra", name: "Algebra liniowa II - Zadania", author: "T. Jurlewicz, Z. Skoczylas", source: "https://docer.pl/doc/esvs1"},
	{category: 0, subcategory: "Algebra", name: "Elementy algebry liniowej I", author: "G. Banaszak, W. Gajda", source: "https://docer.pl/doc/xx8cex8"},
	{category: 0, subcategory: "Algebra", name: "Elementy algebry liniowej II", author: "G. Banaszak, W. Gajda", source: "https://docer.pl/doc/x1vexsv"},
	{category: 0, subcategory: "Algebra", name: "Geometria z algebrą liniową", author: "J. Chaber, R.Pol", source: "https://dydmat.mimuw.edu.pl/sites/default/files/wyklady/geometria-z-algebra-liniowa.pdf"},
	{category: 0, subcategory: "Algebra", name: "Wstęp do algebry - Podstawy algebry", author: "A. Kostrikin", source: "https://docer.pl/doc/xx80nvc"},
	{category: 0, subcategory: "Algebra", name: "Algebra 1", author: "A. Bojanowska, P. Traczyk, ", source: "https://www.mimuw.edu.pl/~wkrych/algebra1-bojanowska-traczyk-stary-program.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna I - Twierdzenia", author: "M. Gewert, Z. Skoczylas", source: "https://lucc.pl/inf/analiza_1/gewert_skoczylas__analiza_matematyczna_1__definicje_twierdzenia_wzory.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna I - Zadania", author: "M. Gewert, Z. Skoczylas", source: "http://staff.uz.zgora.pl/jskowron/pliki/MPT3.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna II - Twierdzenia", author: "M. Gewert, Z. Skoczylas", source: "https://docer.pl/doc/x8e0e0c"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna II - Zadania", author: "M. Gewert, Z. Skoczylas", source: "https://drive.google.com/file/d/0B8QHnlSTD9M8UHRSWWc1djZ1MGM/view?resourcekey=0-6HrD7ig9Pk3-BM60nQKhuA"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Podstawy analizy matematycznej", author: "W. Rudin", source: "https://docer.pl/doc/esv5x"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna I", author: "P. Strzelecki", source: "https://www.mimuw.edu.pl/~pawelst/analiza/Analiza_Matematyczna_1/Notatki_itp./Archiwum_files/skryptAM1-2010-11-ver01.010c.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Analiza Matematyczna II", author: "P. Strzelecki", source: "https://www.mimuw.edu.pl/~pawelst/am2/Analiza_Matematyczna_2/Notatki_files/skrypt-amII-05-01-2016.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Rachunek Różniczkowy i Całkowy I", author: "G. M. Fichtenholz", source: "https://docer.pl/doc/esvs0"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Rachunek Różniczkowy i Całkowy II", author: "G. M. Fichtenholz", source: "https://docer.pl/doc/esvsn"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Rachunek Różniczkowy i Całkowy III", author: "G. M. Fichtenholz", source: "https://docer.pl/doc/ncssvn"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Wstęp do analizy zespolonej", author: "B.W. Szabat ", source: "https://docer.pl/doc/ncn1snx"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Równania różniczkowe zwyczajne", author: "F. Przytycki", source: "https://www.impan.pl/~feliksp/skrypt.pdf"},
	{category: 0, subcategory: "Analiza Matematyczna", name: "Równania różniczkowe zwyczajne", author: "A. Palczewski", source: "https://dydmat.mimuw.edu.pl/sites/default/files/wyklady/rrz.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Geometria trójkąta", author: "S. Zetel", source: "https://docer.pl/doc/x5s08n0"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Wybrane zagadnienia z geometrii płaszczyzny", author: "D. Zaremba", source: "https://www.math.uni.wroc.pl/sites/default/files/repetii.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Topologia I", author: "S. Betley", source: "https://duch.mimuw.edu.pl/~betley/wyklad1/topologia.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Topologia Algebraiczna I", author: "S. Jackowski", source: "https://www.mimuw.edu.pl/~sjack/ta/taI_pomoc_stud.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Topologia I", author: "S. Jackowski", source: "https://www.mimuw.edu.pl/~sjack/Topologia_I/topologia_I_full.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Topologia II", author: "S. Jackowski", source: "https://duch.mimuw.edu.pl/~sjack/dydaktyka/toplogia_II.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Wstęp do geometrii różniczkowej", author: "S. Jackowski", source: "https://www.mimuw.edu.pl/~sjack/dydaktyka/WGR2018Z.pdf"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Wstęp do geometrii różniczkowej", author: "C. Bowszyc", source: "https://docer.pl/doc/1ensvs"},
	{category: 0, subcategory: "Geometria i Topologia", name: "Geometria różniczkowa I", author: "K. Grabowska", source: "https://www.fuw.edu.pl/~konieczn/geometry_lecture."},
	{category: 0, subcategory: "Geometria i Topologia", name: "Geometria różniczkowa", author: "A. Białynicki-Birula", source: "https://www.mimuw.edu.pl/~bbirula/matdyd/g_roz99_00/wyk1.pdf"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Dwanaście wykładów z matematyki obliczeniowej", author: "L. Plaskota, ", source: "https://www.mimuw.edu.pl/~leszekp/dydaktyka/textbook.pdf"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Metody Numeryczne", author: "B. Dahlquist , ", source: "https://docer.pl/doc/s1evs00"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Matematyka Dyskretna", author: "K. Ross", source: "https://docer.pl/doc/s0nvcv"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Matematyka dyskretna dla studentów kierunku Informatyka", author: "H. Furmańczyk", source: "https://inf.ug.edu.pl/~hanna/md/skrypt_okl_full.pdf"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Rachunek Prawdopodobieństwa I", author: "A. Osękowski", source: "https://dydmat.mimuw.edu.pl/nowosci/rachunek-prawdopodobienstwa-i"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Wstęp do rachunku prawdopodobieństwa I", author: "W. Feller", source: "https://docer.pl/doc/ne8s05"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Statystyka - wykłady", author: "A. Krawiec", source: "https://docer.pl/doc/x0cxsv8"},
	{category: 0, subcategory: "Matematyka Dyskretna", name: "Statystyka dla studentów kierunków technicznych i przyrodniczych", author: "J. Koronawski", source: "https://docer.pl/doc/nvenccs"},

	// Biologia
	{category: 1, subcategory: "", name: "Biologia Campbella", author: "N. Campbell", source: "https://docer.pl/doc/x0e15cn"},
	{category: 1, subcategory: "", name: "Biologia - jendość i różnorodność", author: "Praca zbiorowa", source: "https://docer.pl/doc/x1exn1v"},
	{category: 1, subcategory: "Mikrobiologia", name: "Podstawy biologii komorki", author: "B. Alberts, D. Bray, A. Johnson", source: "https://docer.pl/doc/x0svsnx"},
	{category: 1, subcategory: "Mikrobiologia", name: "Podstawy biologii molekularnej", author: "L. Allison", source: "https://docer.pl/doc/x1088n5"},
	{category: 1, subcategory: "Mikrobiologia", name: "Biologia molekularna bakterii", author: "J. Baj", source: "https://docer.pl/doc/ncc88xc"},
	{category: 1, subcategory: "Mikrobiologia", name: "Życie bakrerii", author: "W. Kunicki-Goldfinger", source: "https://docer.pl/doc/vns0xn"},
	{category: 1, subcategory: "Mikrobiologia", name: "Mikrobiologia ogólna", author: "H. G. Schlegel", source: "https://docer.pl/doc/v888es"},
	{category: 1, subcategory: "Mikrobiologia", name: "Krótkie wykłady z mikrobiologii", author: "J. Nicklin", source: "https://docer.pl/doc/x885vn8"},
	{category: 1, subcategory: "Botanika", name: "Botanika I", author: "A. ,J. Szweykowscy", source: "https://docer.pl/doc/xcnense"},
	{category: 1, subcategory: "Botanika", name: "Botanika II", author: "A. ,J. Szweykowscy", source: "https://docer.pl/doc/xcnens8"},
	{category: 1, subcategory: "Botanika", name: "Fizjologia roślin", author: "J. Kopcewicz", source: "https://docer.pl/doc/ses88sn"},
	{category: 1, subcategory: "Zoologia", name: "Zoologia", author: "S. Pisulewski", source: "http://rcin.org.pl/Content/31616/WA058_51142_K856-I_Zoologia-Pisulewski.pdf"},
	{category: 1, subcategory: "Zoologia", name: "Zoologia", author: "W. Lewiński", source: "https://docer.pl/doc/xxx051e"},
	{category: 1, subcategory: "Zoologia", name: "Zoologia", author: "J. Grzegorek", source: "https://docer.pl/doc/n0c81xx"},
	{category: 1, subcategory: "Zoologia", name: "Fizjologia zwierząt", author: "T. Krzymowski", source: "https://docer.pl/doc/v8es1e"},
	{category: 1, subcategory: "Anatomia", name: "Podstawy anatomii człowieka", author: "B. Gołąb", source: "https://docer.pl/doc/xsc10e1"},
	{category: 1, subcategory: "Anatomia", name: "Anatomia i fizjologia człowieka", author: "W. Sylwanowicz", source: "https://docer.pl/doc/xvnn55s"},
	{category: 1, subcategory: "Anatomia", name: "Podstawy immunologii", author: "W. Ptak", source: "https://docer.pl/doc/n0cvs1e"},
	{category: 1, subcategory: "Anatomia", name: "Kompendium Neurologii", author: "R. Podemski", source: "https://docer.pl/doc/x0svses"},
	{category: 1, subcategory: "Anatomia", name: "Krótkie wykłady z neurolobiologii", author: "A. Longstaff", source: "https://docer.pl/doc/x558n10"},
	{category: 1, subcategory: "Biotechnologia", name: "Genetyka molekularna", author: "P. Węgleński ", source: "https://docer.pl/doc/1015es"},
	{category: 1, subcategory: "Biotechnologia", name: "Genomy", author: "T. Brown", source: "https://docer.pl/doc/x01nv5c"},
	{category: 1, subcategory: "Biotechnologia", name: "Podstawy mikrobiologii przemysłowej", author: "J. Paluch", source: "https://docer.pl/doc/s155x0x"},
	{category: 1, subcategory: "Biotechnologia", name: "Bakterie w biologii, biotechnologii i medycynie", author: "P. Singleton", source: "https://docer.pl/doc/xsn1n0s"},
	{category: 1, subcategory: "Biotechnologia", name: "Biofizyka dla studentów", author: "F. Jaroszyk", source: "https://docer.pl/doc/nxsc8"},
	{category: 1, subcategory: "Ekologia i Ewolucja", name: "Krótkie wykłady z ekologii", author: "A. Mackenzie, A. Ball, S. Virdee", source: "https://docer.pl/doc/xncxves"},

	// Chemia
	{category: 2, subcategory: "Chemia Ogólna", name: "Chemia ogólna", author: "L. Jones, P. Atkins", source: "https://docer.pl/doc/nvv1xs0"},
	{category: 2, subcategory: "Chemia Ogólna", name: "Podstawy chemii nieorganicznej I-III", author: "A. Bielański", source: "https://docer.pl/doc/nc0e00c"},
	{category: 2, subcategory: "Chemia Fizyczna", name: "Chemia fizyczna", author: "P. Atkins", source: "https://docer.pl/doc/n10ss5s"},
	{category: 2, subcategory: "Chemia Fizyczna", name: "Elektrochemia I", author: "A. Kisza", source: "https://docer.pl/doc/xs0n8n8"},
	{category: 2, subcategory: "Chemia Fizyczna", name: "Elektrochemia II", author: "A. Kisza", source: "https://docer.pl/doc/xs0n8ne"},
	{category: 2, subcategory: "Chemia Fizyczna", name: "Chemia kwantowa - proste modele", author: "Uniwersytet Warszawski", source: "https://depot.ceon.pl/bitstream/handle/123456789/5367/Chemia_kwantowa_proste_modele_skrypt.pdf?sequence=1&isAllowed=y"},
	{category: 2, subcategory: "Chemia Fizyczna", name: "Krystalografia", author: "Z. Bojarski", source: "https://docer.pl/doc/xcce5ve"},
	{category: 2, subcategory: "Chemia Analityczna", name: "Chemia analityczna I", author: "Minczewski J., Marczenko Z.", source: "https://docer.pl/doc/x00vns5"},
	{category: 2, subcategory: "Chemia Analityczna", name: "Chemia analityczna II", author: "Minczewski J., Marczenko Z.", source: "https://docer.pl/doc/xe1ncxe"},
	{category: 2, subcategory: "Chemia Analityczna", name: "Podstawy chemii analitycznej", author: "Politechnika Gdańska", source: "https://chem.pg.edu.pl/documents/175289/4236621/1Chemia_analityczna_wstep.pdf"},
	{category: 2, subcategory: "Chemia Organiczna", name: "Chemia organiczna", author: "J. McMurry", source: "https://docer.pl/doc/xnnc00x"},
	{category: 2, subcategory: "Chemia Organiczna", name: "Biochemia", author: "J. M. Berg", source: "https://docer.pl/doc/xncvec8"},

	// Fizyka
	{category: 3, subcategory: "", name: "Fizyka Doświadczalna", author: "H. Szydlowski", source: "https://docer.pl/doc/n0n8cxn"},
	{category: 3, subcategory: "Mechanika", name: "Podstawy Fizyki I", author: "D. Halliday", source: "https://docer.pl/doc/n5508c1"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika", author: "L. Landau", source: "https://docer.pl/doc/xexvn58"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika", author: "C. Kittel", source: "https://docer.pl/doc/n0sxxx0"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika ogólna I", author: "J. Leyko", source: "https://docer.pl/doc/x0xs108"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika ogólna II", author: "J. Leyko", source: "https://docer.pl/doc/xnc0s0c"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika ogólna I", author: "J. Misiak ", source: "https://docer.pl/doc/n1css1e"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika ogólna II", author: "J. Misiak ", source: "https://docer.pl/doc/vn8801"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika klasyczna I", author: "J. Taylor", source: "https://docer.pl/doc/env0vvv"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika klasyczna II", author: "J. Taylor", source: "https://docer.pl/doc/env0vvc"},
	{category: 3, subcategory: "Mechanika", name: "Mechanika techniczna", author: "W. Siuta", source: "https://docer.pl/doc/x0en5vs"},
	{category: 3, subcategory: "Mechanika", name: "Fale", author: "F. Crawford", source: "https://docer.pl/doc/xe81ecv"},
	{category: 3, subcategory: "Mechanika", name: "Fizyka I", author: "Openstax", source: "https://openstax.org/details/books/fizyka-dla-szk%C3%B3%C5%82-wy%C5%BCszych-tom-1"},
	{category: 3, subcategory: "Termodynamika", name: "Podstawy Fizyki II", author: "D. Halliday", source: "https://docer.pl/doc/n5c00vn"},
	{category: 3, subcategory: "Termodynamika", name: "Termodynamika", author: "W. Pudlik", source: "https://pbc.gda.pl/Content/4098/pbc_termodynamika.pdf"},
	{category: 3, subcategory: "Termodynamika", name: "Termodynamika", author: "B. Staniszewski", source: "https://docer.pl/doc/558s0"},
	{category: 3, subcategory: "Termodynamika", name: "Termodynamika techniczna", author: "J. Szargut", source: "https://docer.pl/doc/nvce8sn"},
	{category: 3, subcategory: "Termodynamika", name: "Termodynamika dla fizyków, chemików i inżynierów", author: "R. Hołyst", source: "https://docer.pl/doc/x18xxs0"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Podstawy Fizyki III", author: "D. Halliday", source: "https://docer.pl/doc/xeenece"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Elektryczność i magnetyzm", author: "A. Piekara", source: "https://docer.pl/doc/xce8ssx"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Elektryczność i magnetyzm", author: "E. Purcell", source: "https://docer.pl/doc/xc18ve"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Elektryczność i magnetyzm", author: "P. Kossacki", source: "http://lumnp.fuw.edu.pl/kossacki/elektrycznoscimagnetyzm/"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Podstawy elektrodynamiki", author: "D. Griffiths", source: "https://docer.pl/doc/n8sx5ne"},
	{category: 3, subcategory: "Elektromagnetyzm", name: "Fizyka II", author: "Openstax", source: "https://openstax.org/details/books/fizyka-dla-szk%C3%B3%C5%82-wy%C5%BCszych-tom-2"},
	{category: 3, subcategory: "Optyka", name: "Podstawy Fizyki IV", author: "D. Halliday", source: "https://docer.pl/doc/xeenecs"},
	{category: 3, subcategory: "Optyka", name: "Optyka - kurs elementarny", author: "J. Nowak", source: "https://docer.pl/doc/sx1vxs8"},
	{category: 3, subcategory: "Optyka", name: "Fizyka laserów", author: "B. Ziętek", source: "https://fizyka.umk.pl/~bezet/pdf/II%20Fizyka%20laserow.pdf"},
	{category: 3, subcategory: "Fizyka Kwantowa", name: "Mechanika Kwantowa", author: "L. Landau", source: "https://docer.pl/doc/xexvn58"},
	{category: 3, subcategory: "Fizyka Kwantowa", name: "Mechanika Kwantowa", author: "S. Kryszewski", source: "https://www.fuw.edu.pl/~matri/mechemq/QM_SKryszewski.pdf"},
	{category: 3, subcategory: "Fizyka Kwantowa", name: "Wykłady z mechaniki kwantowej", author: "R. Feynman", source: "https://docer.pl/doc/xe81ec1"},
	{category: 3, subcategory: "Fizyka Kwantowa", name: "Wstęp do informatyki kwantowej", author: "M. Le Bellac", source: "https://docer.pl/doc/sexncx1"},
	{category: 3, subcategory: "Fizyka Kwantowa", name: "Kondensat Bosego-Einsteina", author: "J. Matulewski", source: "https://fizyka.umk.pl/~jacek/dydaktyka/mkdn/07_KS_BEC.pdf"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Podstawy Fizyki V", author: "D. Halliday", source: "https://docer.pl/doc/xeenecn"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Wstęp do fizyki atomu, cząsteczki i ciała stałego", author: "J. Ginter", source: "https://docer.pl/doc/xe1nc05"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Encyklopedia fizyki współczesnej", author: "A. Wróblewski", source: "https://docer.pl/doc/x1818"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Atom i kwanty", author: "H. Haken", source: "https://docer.pl/doc/xcs8nv"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Fizyka molekularna z elementami chemii kwantowej", author: "H. Hermann", source: "https://docer.pl/doc/ne5v10e"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Wprowadzenie do fizyki subatomowej", author: "M. Pfützner", source: "https://www.fuw.edu.pl/~pfutzner/Teaching/Wyklady/WprowadzenieDoFizykiSubatomowej.pdf"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Wybrane zagadnienia fizyki subatomowej", author: "Z. Janas", source: "https://www.fuw.edu.pl/~szczytko/FMS/Fiz_Subatom_IN.pdf"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Fizyka ciała stałego", author: "Ch. Kittel", source: "https://docer.pl/doc/x0xneve"},
	{category: 3, subcategory: "Fizyka Atomowa", name: "Fizyka III", author: "Openstax", source: "https://openstax.org/details/books/fizyka-dla-szk%C3%B3%C5%82-wy%C5%BCszych-tom-3"},

	// Geografia
	{category: 4, subcategory: "Kartografia", name: "Wprowadzenie do kartografii i topografii ", author: "J. Pasławski", source: "https://docer.pl/doc/sxcv505"},
	{category: 4, subcategory: "Meteorologia", name: "Podstawy meteorologii i klimatologii Polski", author: "C. Walczakiewicz", source: "https://www.researchgate.net/publication/332345347_Podstawy_meteorologii_i_klimatologii_Polski"},
	{category: 4, subcategory: "Meteorologia", name: "Meteorologia i klimatologia", author: "K. Kożuchowski", source: "https://docer.pl/doc/xs1esv1"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Geografia fizyczna Polski", author: "A. Richling", source: "https://docer.pl/doc/svc8xvn"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Hydrologia ogólna", author: "E. Bajkiewicz-Grabowska", source: "https://docer.pl/doc/xnnn55s"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Geomorfologia", author: "M. Klimaszewski", source: "https://docer.pl/doc/xc18en1"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Ekologia z biogeografią i ochroną środowiska", author: "B. Bukała", source: "https://docer.pl/doc/v58s85"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Historia ziemii", author: "S. Stanley", source: "https://docer.pl/doc/x1558s8"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Systematyka gleb Polski", author: "M. Świtoniak", source: "https://docer.pl/doc/s815v8n"},
	{category: 4, subcategory: "Geografia Fizyczna", name: "Gleboznawstwo", author: "S. Zawadzki", source: "https://docer.pl/doc/xs1es58"},
	{category: 4, subcategory: "Geografia Społeczna", name: "Geografia polityczna ogólna", author: "J. Barbag", source: "https://docer.pl/doc/e0ssn"},
	{category: 4, subcategory: "Geografia Społeczna", name: "Geografia ekonomiczna - Ujęcie dynamiczne", author: "R. Domański", source: "https://docer.pl/doc/n1snssc"},
	{category: 4, subcategory: "Geografia Społeczna", name: "Geografia ekonomiczna. Zarys teoretyczny", author: "K. Kuciński", source: "https://docer.pl/doc/s155e85"},

	// // Informatyka
	// {category: 6, subcategory: "Programowanie", name: "C# - Praktyczny kurs", author: "M. Lis", source: "52270699"},
	// {category: 6, subcategory: "Programowanie", name: "Szkoła programowania w C++", author: "S. Prata", source: "39331449"},
	// {category: 6, subcategory: "Programowanie", name: "Język Go - poznaj i programuj", author: "A. Donovan", source: "37209230"},
	// {category: 6, subcategory: "Programowanie", name: "Python - programuj szybko i wydajnie", author: "M. Goelick", source: "66770043"},
	// {category: 6, subcategory: "Programowanie", name: "Dart i Flutter dla początkujących", author: "A. Biessek", source: "54659425"},

	// ########################################################################
	// ########################################################################
	// ########################################################################
	// ########################################################################

	// Ekonomia
	{category: 6, subcategory: "", name: "Podstawy ekonomii", author: "R. Milewski", source: "https://docer.pl/doc/s1cc5c"},
	{category: 6, subcategory: "", name: "Ekonomia dla prawników i nie tylko", author: "M. Bednarski", source: "https://docer.pl/doc/xsc01vx"},
	{category: 6, subcategory: "", name: "Prawo Handlowe", author: "A. Kidyba", source: "https://docer.pl/doc/xssssxs"},
	{category: 6, subcategory: "Mikroekonomia", name: "Mikroekonomia", author: "N. Gregory", source: "https://docer.pl/doc/nn8evs"},
	{category: 6, subcategory: "Mikroekonomia", name: "Mikroekonomia - kurs średni", author: "H. Varian", source: "https://docer.pl/doc/nn55v1"},
	{category: 6, subcategory: "Mikroekonomia", name: "Podstawy zarządzania organizacjami", author: "R.W. Griffin", source: "https://docer.pl/doc/8veccs"},
	{category: 6, subcategory: "Mikroekonomia", name: "Współczesne teorie przedsiębiorstwa", author: "T. Gruszecki", source: "https://docer.pl/doc/x1c085e"},
	{category: 6, subcategory: "Mikroekonomia", name: "Advanced international trade: theory and evsourceence", author: "R. Feenstra", source: "http://course.sdu.edu.cn/G2S/eWebEditor/uploadfile/20120417190154_999045912265.pdf"},
	{category: 6, subcategory: "Mikroekonomia", name: "Ekonomia cz. I", author: "P. Samuelson", source: "https://docer.pl/doc/n5505se"},
	{category: 6, subcategory: "Makroekonomia", name: "Makroekonomia", author: "N. Mankiw", source: "https://docer.pl/doc/ncc88ne"},
	{category: 6, subcategory: "Makroekonomia", name: "Makroekonomia", author: "D. Begg", source: "https://docer.pl/doc/x1s5xce"},
	{category: 6, subcategory: "Makroekonomia", name: "Makroekonomia", author: "O. Blanchard", source: "https://docer.pl/doc/x8085s8"},
	{category: 6, subcategory: "Mikroekonomia", name: "Ekonomia cz. II", author: "P. Samuelson", source: "https://docer.pl/doc/n5505sv"},
	{category: 6, subcategory: "Polityka Gospodarcza", name: "Polityka gospodarcza: wyzwania, dylematy, priorytety", author: "J. Stacewicz", source: "https://ssl-kolegia.sgh.waw.pl/pl/KAE/struktura/IRG/publikacje/Documents/pim83.pdf"},
	{category: 6, subcategory: "Polityka Gospodarcza", name: "Współczesna gospodarka światowa", author: "A. Kisiel-Łowczyc", source: "https://docer.pl/doc/x8svvnx"},
	{category: 6, subcategory: "Ekonometria", name: "Ekonometria I", author: "M. Rubaszek", source: "https://web.sgh.waw.pl/~mrubas/Econometrics/pdf/EI_TallPL.pdf"},
	{category: 6, subcategory: "Ekonometria", name: "Ekonometria II", author: "M. Rubaszek", source: "https://web.sgh.waw.pl/~mrubas/EFII/pdf/SkryptEF2.pdf"},
	{category: 6, subcategory: "Ekonometria", name: "Informatyka ekonomiczna", author: "S. Wrycza", source: "https://docer.pl/doc/xs8exx1"},
	{category: 6, subcategory: "Ekonometria", name: "Podstawy ekonometrii z elementami algebry liniowej", author: "W. Nowakowski", source: "https://wszechnicapolska.edu.pl/dokumenty/wydawnictwo/2011-E-W-Nowakowski-Podstawy-ekonometrii-z-elementami-algebry-liniowej.pdf"},
	{category: 6, subcategory: "Ekonometria", name: "Modelowanie ekonometryczne", author: "B. Gładysz", source: "https://dbc.wroc.pl/Content/2182/PDF/Gladysz_modelowanie_ekonometryczne.pdf"},
	{category: 6, subcategory: "Ekonometria", name: "Systemy informatyczne zarządzania przedsiębiorstwem", author: "Z. Klonowski", source: "https://www.dbc.wroc.pl/Content/968/klonowski_systemy_informatyczne.pdf"},
	{category: 6, subcategory: "Ekonometria", name: "Badania operacyjne w przykladach i zadaniach ", author: "Z. Jedrzejczyk", source: "https://docer.pl/doc/n5ve8v1"},
	{category: 6, subcategory: "Ekonometria", name: "Wprowadzenie do Badań Operacyjnych z Komputerem", author: "T.Trzaskalik", source: "https://docer.pl/doc/xxves00"},
	{category: 6, subcategory: "Finanse", name: "Finanse", author: "S. Owsiak", source: "https://docer.pl/doc/n00xx0x"},
	{category: 6, subcategory: "Finanse", name: "Finanse publiczne - teoria i praktyka", author: "S. Owsiak", source: "https://docer.pl/doc/ccn85c"},
	{category: 6, subcategory: "Finanse", name: "Inwestycje i instrumenty finansowe", author: "K. & T. Jajuga", source: "https://docer.pl/doc/xxvn11x"},
	{category: 6, subcategory: "Finanse", name: "Finanse publiczne i prawo finansowe - realia i perspektywy zmian", author: "L. Etel", source: "https://repozytorium.uwb.edu.pl/jspui/bitstream/11320/5923/1/Finanse_publiczne_-_realia_i_perspektywy_zmian.pdf"},
	{category: 6, subcategory: "Finanse", name: "Zarządzanie finansami I", author: "E. Brigham", source: "https://docer.pl/doc/nev0815"},
	{category: 6, subcategory: "Finanse", name: "Zarządzanie finansami II", author: "E. Brigham", source: "https://docer.pl/doc/nev08s1"},
	{category: 6, subcategory: "Finanse", name: "System podatkowy w Polsce", author: "R. Wolański", source: "https://docer.pl/doc/xex81nx"},
	{category: 6, subcategory: "Finanse", name: "Analiza finansowa przedsiębiorstwa - teoria i praktyka", author: "W. Gabrusewicz", source: "https://docer.pl/doc/xn1cvev"},
	{category: 6, subcategory: "Finanse", name: "Zarządzanie finansami przedsiębiorstwa", author: "W. Bień", source: "https://docer.pl/doc/x1e5n1s"},
	{category: 6, subcategory: "Finanse", name: "Zarządzanie finansami przedsiębiorstw - podstawy teorii", author: "J. Czekaj", source: "https://docer.pl/doc/x0n1vv8"},
	{category: 6, subcategory: "Finanse", name: "Zarządzanie - teoria i praktyka", author: "A. Koźmiński", source: "https://docer.pl/doc/505n"},
	{category: 6, subcategory: "Rachunkowość", name: "Rachunkowość finansowa I", author: "Z. Messner", source: "https://docer.pl/doc/nevn50v"},
	{category: 6, subcategory: "Rachunkowość", name: "Rachunkowość finansowa II", author: "Z. Messner", source: "https://docer.pl/doc/n8vsnxx"},
	{category: 6, subcategory: "Rachunkowość", name: "Rachunkowość finansowa z uwzględnieniem MSSF", author: "Z. Messner", source: "https://docer.pl/doc/n00n5n5"},
	{category: 6, subcategory: "Rachunkowość", name: "Rachunkowość od podstaw", author: "D. Dobija", source: "https://docer.pl/doc/nevn81v"},
	{category: 6, subcategory: "Rachunkowość", name: "Rachunkowość finansowa w teorii i praktyce", author: "E. Kalwasińska", source: "https://docer.pl/doc/xnsnnv8"},

	// ########################################################################
	// ########################################################################
	// ########################################################################
	// ########################################################################

	// Filozofia
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii I", author: "W. Tatarkiewicz", source: "https://docer.pl/doc/s110ee5"},
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii II", author: "W. Tatarkiewicz", source: "https://docer.pl/doc/s110eee"},
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii III", author: "W. Tatarkiewicz", source: "https://docer.pl/doc/s110ee8"},
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii starożytnej I", author: "G. Reale", source: "https://docer.pl/doc/5xxesc"},
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii starożytnej II", author: "G. Reale", source: "https://docer.pl/doc/nc18sxn"},
	{category: 20, subcategory: "Historia Filozofii", name: "Historia filozofii starożytnej III", author: "G. Reale", source: "https://docer.pl/doc/xecxcnn"},
	{category: 20, subcategory: "Filozofia", name: "Filozofia dla bystrzaków", author: "T. Morris", source: "https://docer.pl/doc/nsnvexs"},
	{category: 20, subcategory: "Filozofia", name: "Zagadnienia i kierunki filozofii - teoria poznania", author: "K. Ajdukiewicz", source: "https://docer.pl/doc/nvn5s0x"},
	{category: 20, subcategory: "Filozofia", name: "Estetyka i antyestetyka", author: "M. Gołaszewska", source: "https://docer.pl/doc/nxev81"},
	{category: 20, subcategory: "Logika", name: "Logika dla bystrzaków", author: "M. Zegarelli", source: "https://docer.pl/doc/enxcxcx"},
	{category: 20, subcategory: "Logika", name: "Zarys logiki", author: "M. Omyła", source: "https://docer.pl/doc/xcve808"},
	{category: 20, subcategory: "Logika", name: "Elementarne wprowadzenie w logikę formalną", author: "M. Porębska", source: "https://docer.pl/doc/xvvsnce"},

	// Historia
	{category: 21, subcategory: "Starożytność", name: "Historia Starożytna", author: "M. Jaczynowska", source: "https://docer.pl/doc/xx0s8es"},
	{category: 21, subcategory: "Starożytność", name: "Starożytny Rzym", author: "M. Jaczynowska", source: "https://docer.pl/doc/xc0c8ec"},
	{category: 21, subcategory: "Starożytność", name: "Od Homera do Kleopatry", author: "D. Musiał", source: "https://docer.pl/doc/x0nx0n"},
	{category: 21, subcategory: "Starożytność", name: "Starożytna Grecja okresu archaicznego i klasycznego", author: "W. Lengauer", source: "https://docer.pl/doc/xsvx1ns"},

	{category: 21, subcategory: "Średniowiecze", name: "Historia kultury bizantyńskiej", author: "H. Haussig", source: "https://docer.pl/doc/xesnx08"},
	{category: 21, subcategory: "Średniowiecze", name: "Cesarstwo Bizantyńskie", author: "R. Browning", source: "https://docer.pl/doc/nn5vnx"},
	{category: 21, subcategory: "Średniowiecze", name: "Historia powszechna, średniowiecze", author: "T. Manteuffel", source: "https://docer.pl/doc/x0cxe58"},
	{category: 21, subcategory: "Średniowiecze", name: "Historia powszechna, średniowiecze", author: "R. Michałowski", source: "https://docer.pl/doc/n80050n"},

	{category: 21, subcategory: "Nowożytność", name: "", author: "", source: ""},
	{category: 21, subcategory: "Współczesność", name: "", author: "", source: ""},

	// // Psychologia
	// {category: 9, subcategory: "", name: "Psychologia - Kluczowe koncepcje tom I", author: "P. Zimbardo", source: "66648907"},
	// {category: 9, subcategory: "", name: "Psychologia - Kluczowe koncepcje tom II", author: "P. Zimbardo", source: "51936913"},
	// {category: 9, subcategory: "", name: "Psychologia - Kluczowe koncepcje tom III", author: "P. Zimbardo", source: "62329168"},
	// {category: 9, subcategory: "", name: "Psychologia - Kluczowe koncepcje tom IV", author: "P. Zimbardo", source: "78338516"},
	// {category: 9, subcategory: "", name: "Psychologia - Kluczowe koncepcje tom V", author: "P. Zimbardo", source: "43656557"},

	// // ########################################################################
	// // ########################################################################
	// // ########################################################################
	// // ########################################################################

	// // Elektronika
	// {category: 10, subcategory: "", name: "Podstawy Elektroniki", author: "???", source: "61977616"},
	// {category: 10, subcategory: "", name: "Podstawy Elektroniki", author: "Politechnika Wrocławska", source: "33690110"},
	// {category: 10, subcategory: "", name: "Sztuka Elektroniki 2 Tomy", author: "P. Horowitz", source: "74992378"},

	// // Geologia
	// {category: 5, subcategory: "", name: "", author: "", source: ""},

	// // ########################################################################
	// // ########################################################################
	// // ########################################################################
	// // ########################################################################

	// {category: len(categories) - 1, subcategory: "", name: "Lemegeton", author: "Salomon", source: "30651537"},
	// {category: len(categories) - 1, subcategory: "", name: "Mein Kampf", author: "A. Hitler", source: "23382005"},
}
