const relaysSupportedData = [
  {
    value: 'Aestus',
    label: 'Aestus',
  },

  {
    value: 'Agnostic Gnosis',
    label: 'Agnostic Gnosis',
  },
  {
    value: 'Blocknative',
    label: 'Blocknative',
  },

  {
    value: 'bloXroute',
    label: 'bloXroute',
  },
  {
    value: 'Eden Network',
    label: 'Eden Network',
  },

  {
    value: 'Flashbots',
    label: 'Flashbots',
  },
  {
    value: 'Manifold',
    label: 'Manifold',
  },

  {
    value: 'Ultra Sound',
    label: 'Ultra Sound',
  },
];

const locationData = [
  {
    label: 'Afghanistan (افغانستان‎)',
    value: 'Afghanistan (افغانستان‎)|af',
    abbr: 'af',
  },
  {
    label: 'Albania (Shqipëri)',
    value: 'Albania (Shqipëri)|al',
  },
  {
    label: 'Algeria (الجزائر‎)',
    value: 'Algeria (الجزائر‎)|dz',
  },
  {
    label: 'American Samoa',
    value: 'American Samoa|as',
  },
  {
    label: 'Andorra',
    value: 'Andorra|ad',
  },
  {
    label: 'Angola',
    value: 'Angola|ao',
  },
  {
    label: 'Anguilla',
    value: 'Anguilla|ai',
  },
  {
    label: 'Antigua and Barbuda',
    value: 'Antigua and Barbuda|ag',
  },
  {
    label: 'Argentina',
    value: 'Argentina|ar',
  },
  {
    label: 'Armenia (Հայաստան)',
    value: 'Armenia (Հայաստան)|am',
  },
  {
    label: 'Aruba',
    value: 'Aruba|aw',
  },
  {
    label: 'Australia',
    value: 'Australia|au',
  },
  {
    label: 'Austria (Österreich)',
    value: 'Austria (Österreich)|at',
  },
  {
    label: 'Azerbaijan (Azərbaycan)',
    value: 'Azerbaijan (Azərbaycan)|az',
  },
  {
    label: 'Bahamas',
    value: 'Bahamas|bs',
  },
  {
    label: 'Bahrain (البحرين‎)',
    value: 'Bahrain (البحرين‎)|bh',
  },
  {
    label: 'Bangladesh (বাংলাদেশ)',
    value: 'Bangladesh (বাংলাদেশ)|bd',
  },
  {
    label: 'Barbados',
    value: 'Barbados|bb',
  },
  {
    label: 'Belarus (Беларусь)',
    value: 'Belarus (Беларусь)|by',
  },
  {
    label: 'Belgium (België)',
    value: 'Belgium (België)|be',
  },
  {
    label: 'Belize',
    value: 'Belize|bz',
  },
  {
    label: 'Benin (Bénin)',
    value: 'Benin (Bénin)|bj',
  },
  {
    label: 'Bermuda',
    value: 'Bermuda|bm',
  },
  {
    label: 'Bhutan (འབྲུག)',
    value: 'Bhutan (འབྲུག)|bt',
  },
  {
    label: 'Bolivia',
    value: 'Bolivia|bo',
  },
  {
    label: 'Bosnia and Herzegovina (Босна и Херцеговина)',
    value: 'Bosnia and Herzegovina (Босна и Херцеговина)|ba',
  },
  {
    label: 'Botswana',
    value: 'Botswana|bw',
  },
  {
    label: 'Brazil (Brasil)',
    value: 'Brazil (Brasil)|br',
  },
  {
    label: 'British Indian Ocean Territory',
    value: 'British Indian Ocean Territory|io',
  },
  {
    label: 'British Virgin Islands',
    value: 'British Virgin Islands|vg',
  },
  {
    label: 'Brunei',
    value: 'Brunei|bn',
  },
  {
    label: 'Bulgaria (България)',
    value: 'Bulgaria (България)|bg',
  },
  {
    label: 'Burkina Faso',
    value: 'Burkina Faso|bf',
  },
  {
    label: 'Burundi (Uburundi)',
    value: 'Burundi (Uburundi)|bi',
  },
  {
    label: 'Cambodia (កម្ពុជា)',
    value: 'Cambodia (កម្ពុជា)|kh',
  },
  {
    label: 'Cameroon (Cameroun)',
    value: 'Cameroon (Cameroun)|cm',
  },
  {
    label: 'Canada',
    value: 'Canada|ca',
  },
  {
    label: 'Cape Verde (Kabu Verdi)',
    value: 'Cape Verde (Kabu Verdi)|cv',
  },
  {
    label: 'Caribbean Netherlands',
    value: 'Caribbean Netherlands|bq',
  },
  {
    label: 'Cayman Islands',
    value: 'Cayman Islands|ky',
  },
  {
    label: 'Central African Republic (République centrafricaine)',
    value: 'Central African Republic (République centrafricaine)|cf',
  },
  {
    label: 'Chad (Tchad)',
    value: 'Chad (Tchad)|td',
  },
  {
    label: 'Chile',
    value: 'Chile|cl',
  },
  {
    label: 'China (中国)',
    value: 'China (中国)|cn',
  },
  {
    label: 'Christmas Island',
    value: 'Christmas Island|cx',
  },
  {
    label: 'Cocos (Keeling) Islands',
    value: 'Cocos (Keeling) Islands|cc',
  },
  {
    label: 'Colombia',
    value: 'Colombia|co',
  },
  {
    label: 'Comoros (جزر القمر‎)',
    value: 'Comoros (جزر القمر‎)|km',
  },
  {
    label: 'Congo (DRC) (Jamhuri ya Kidemokrasia ya Kongo)',
    value: 'Congo (DRC) (Jamhuri ya Kidemokrasia ya Kongo)|cd',
  },
  {
    label: 'Congo (Republic) (Congo-Brazzaville)',
    value: 'Congo (Republic) (Congo-Brazzaville)|cg',
  },
  {
    label: 'Cook Islands',
    value: 'Cook Islands|ck',
  },
  {
    label: 'Costa Rica',
    value: 'Costa Rica|cr',
  },
  {
    label: 'Côte d’Ivoire',
    value: 'Côte d’Ivoire|ci',
  },
  {
    label: 'Croatia (Hrvatska)',
    value: 'Croatia (Hrvatska)|hr',
  },
  {
    label: 'Cuba',
    value: 'Cuba|cu',
  },
  {
    label: 'Curaçao',
    value: 'Curaçao|cw',
  },
  {
    label: 'Cyprus (Κύπρος)',
    value: 'Cyprus (Κύπρος)|cy',
  },
  {
    label: 'Czech Republic (Česká republika)',
    value: 'Czech Republic (Česká republika)|cz',
  },
  {
    label: 'Denmark (Danmark)',
    value: 'Denmark (Danmark)|dk',
  },
  {
    label: 'Djibouti',
    value: 'Djibouti|dj',
  },
  {
    label: 'Dominica',
    value: 'Dominica|dm',
  },
  {
    label: 'Dominican Republic (República Dominicana)',
    value: 'Dominican Republic (República Dominicana)|do',
  },
  {
    label: 'Ecuador',
    value: 'Ecuador|ec',
  },
  {
    label: 'Egypt (مصر‎)',
    value: 'Egypt (مصر‎)|eg',
  },
  {
    label: 'El Salvador',
    value: 'El Salvador|sv',
  },
  {
    label: 'Equatorial Guinea (Guinea Ecuatorial)',
    value: 'Equatorial Guinea (Guinea Ecuatorial)|gq',
  },
  {
    label: 'Eritrea',
    value: 'Eritrea|er',
  },
  {
    label: 'Estonia (Eesti)',
    value: 'Estonia (Eesti)|ee',
  },
  {
    label: 'Ethiopia',
    value: 'Ethiopia|et',
  },
  {
    label: 'Falkland Islands (Islas Malvinas)',
    value: 'Falkland Islands (Islas Malvinas)|fk',
  },
  {
    label: 'Faroe Islands (Føroyar)',
    value: 'Faroe Islands (Føroyar)|fo',
  },
  {
    label: 'Fiji',
    value: 'Fiji|fj',
  },
  {
    label: 'Finland (Suomi)',
    value: 'Finland (Suomi)|fi',
  },
  {
    label: 'France',
    value: 'France|fr',
  },
  {
    label: 'French Guiana (Guyane française)',
    value: 'French Guiana (Guyane française)|gf',
  },
  {
    label: 'French Polynesia (Polynésie française)',
    value: 'French Polynesia (Polynésie française)|pf',
  },
  {
    label: 'Gabon',
    value: 'Gabon|ga',
  },
  {
    label: 'Gambia',
    value: 'Gambia|gm',
  },
  {
    label: 'Georgia (საქართველო)',
    value: 'Georgia (საქართველო)|ge',
  },
  {
    label: 'Germany (Deutschland)',
    value: 'Germany (Deutschland)|de',
  },
  {
    label: 'Ghana (Gaana)',
    value: 'Ghana (Gaana)|gh',
  },
  {
    label: 'Gibraltar',
    value: 'Gibraltar|gi',
  },
  {
    label: 'Greece (Ελλάδα)',
    value: 'Greece (Ελλάδα)|gr',
  },
  {
    label: 'Greenland (Kalaallit Nunaat)',
    value: 'Greenland (Kalaallit Nunaat)|gl',
  },
  {
    label: 'Grenada',
    value: 'Grenada|gd',
  },
  {
    label: 'Guadeloupe',
    value: 'Guadeloupe|gp',
  },
  {
    label: 'Guam',
    value: 'Guam|gu',
  },
  {
    label: 'Guatemala',
    value: 'Guatemala|gt',
  },
  {
    label: 'Guernsey',
    value: 'Guernsey|gg',
  },
  {
    label: 'Guinea (Guinée)',
    value: 'Guinea (Guinée)|gn',
  },
  {
    label: 'Guinea-Bissau (Guiné Bissau)',
    value: 'Guinea-Bissau (Guiné Bissau)|gw',
  },
  {
    label: 'Guyana',
    value: 'Guyana|gy',
  },
  {
    label: 'Haiti',
    value: 'Haiti|ht',
  },
  {
    label: 'Honduras',
    value: 'Honduras|hn',
  },
  {
    label: 'Hong Kong (香港)',
    value: 'Hong Kong (香港)|hk',
  },
  {
    label: 'Hungary (Magyarország)',
    value: 'Hungary (Magyarország)|hu',
  },
  {
    label: 'Iceland (Ísland)',
    value: 'Iceland (Ísland)|is',
  },
  {
    label: 'India (भारत)',
    value: 'India (भारत)|in',
  },
  {
    label: 'Indonesia',
    value: 'Indonesia|id',
  },
  {
    label: 'Iran (ایران‎)',
    value: 'Iran (ایران‎)|ir',
  },
  {
    label: 'Iraq (العراق‎)',
    value: 'Iraq (العراق‎)|iq',
  },
  {
    label: 'Ireland',
    value: 'Ireland|ie',
  },
  {
    label: 'Isle of Man',
    value: 'Isle of Man|im',
  },
  {
    label: 'Israel (ישראל‎)',
    value: 'Israel (ישראל‎)|il',
  },
  {
    label: 'Italy (Italia)',
    value: 'Italy (Italia)|it',
  },
  {
    label: 'Jamaica',
    value: 'Jamaica|jm',
  },
  {
    label: 'Japan (日本)',
    value: 'Japan (日本)|jp',
  },
  {
    label: 'Jersey',
    value: 'Jersey|je',
  },
  {
    label: 'Jordan (الأردن‎)',
    value: 'Jordan (الأردن‎)|jo',
  },
  {
    label: 'Kazakhstan (Казахстан)',
    value: 'Kazakhstan (Казахстан)|kz',
  },
  {
    label: 'Kenya',
    value: 'Kenya|ke',
  },
  {
    label: 'Kiribati',
    value: 'Kiribati|ki',
  },
  {
    label: 'Kosovo',
    value: 'Kosovo|xk',
  },
  {
    label: 'Kuwait (الكويت‎)',
    value: 'Kuwait (الكويت‎)|kw',
  },
  {
    label: 'Kyrgyzstan (Кыргызстан)',
    value: 'Kyrgyzstan (Кыргызстан)|kg',
  },
  {
    label: 'Laos (ລາວ)',
    value: 'Laos (ລາວ)|la',
  },
  {
    label: 'Latvia (Latvija)',
    value: 'Latvia (Latvija)|lv',
  },
  {
    label: 'Lebanon (لبنان‎)',
    value: 'Lebanon (لبنان‎)|lb',
  },
  {
    label: 'Lesotho',
    value: 'Lesotho|ls',
  },
  {
    label: 'Liberia',
    value: 'Liberia|lr',
  },
  {
    label: 'Libya (ليبيا‎)',
    value: 'Libya (ليبيا‎)|ly',
  },
  {
    label: 'Liechtenstein',
    value: 'Liechtenstein|li',
  },
  {
    label: 'Lithuania (Lietuva)',
    value: 'Lithuania (Lietuva)|lt',
  },
  {
    label: 'Luxembourg',
    value: 'Luxembourg|lu',
  },
  {
    label: 'Macau (澳門)',
    value: 'Macau (澳門)|mo',
  },
  {
    label: 'Macedonia (FYROM) (Македонија)',
    value: 'Macedonia (FYROM) (Македонија)|mk',
  },
  {
    label: 'Madagascar (Madagasikara)',
    value: 'Madagascar (Madagasikara)|mg',
  },
  {
    label: 'Malawi',
    value: 'Malawi|mw',
  },
  {
    label: 'Malaysia',
    value: 'Malaysia|my',
  },
  {
    label: 'Maldives',
    value: 'Maldives|mv',
  },
  {
    label: 'Mali',
    value: 'Mali|ml',
  },
  {
    label: 'Malta',
    value: 'Malta|mt',
  },
  {
    label: 'Marshall Islands',
    value: 'Marshall Islands|mh',
  },
  {
    label: 'Martinique',
    value: 'Martinique|mq',
  },
  {
    label: 'Mauritania (موريتانيا‎)',
    value: 'Mauritania (موريتانيا‎)|mr',
  },
  {
    label: 'Mauritius (Moris)',
    value: 'Mauritius (Moris)|mu',
  },
  {
    label: 'Mayotte',
    value: 'Mayotte|yt',
  },
  {
    label: 'Mexico (México)',
    value: 'Mexico (México)|mx',
  },
  {
    label: 'Micronesia',
    value: 'Micronesia|fm',
  },
  {
    label: 'Moldova (Republica Moldova)',
    value: 'Moldova (Republica Moldova)|md',
  },
  {
    label: 'Monaco',
    value: 'Monaco|mc',
  },
  {
    label: 'Mongolia (Монгол)',
    value: 'Mongolia (Монгол)|mn',
  },
  {
    label: 'Montenegro (Crna Gora)',
    value: 'Montenegro (Crna Gora)|me',
  },
  {
    label: 'Montserrat',
    value: 'Montserrat|ms',
  },
  {
    label: 'Morocco (المغرب‎)',
    value: 'Morocco (المغرب‎)|ma',
  },
  {
    label: 'Mozambique (Moçambique)',
    value: 'Mozambique (Moçambique)|mz',
  },
  {
    label: 'Myanmar (Burma) (မြန်မာ)',
    value: 'Myanmar (Burma) (မြန်မာ)|mm',
  },
  {
    label: 'Namibia (Namibië)',
    value: 'Namibia (Namibië)|na',
  },
  {
    label: 'Nauru',
    value: 'Nauru|nr',
  },
  {
    label: 'Nepal (नेपाल)',
    value: 'Nepal (नेपाल)|np',
  },
  {
    label: 'Netherlands (Nederland)',
    value: 'Netherlands (Nederland)|nl',
  },
  {
    label: 'New Caledonia (Nouvelle-Calédonie)',
    value: 'New Caledonia (Nouvelle-Calédonie)|nc',
  },
  {
    label: 'New Zealand',
    value: 'New Zealand|nz',
  },
  {
    label: 'Nicaragua',
    value: 'Nicaragua|ni',
  },
  {
    label: 'Niger (Nijar)',
    value: 'Niger (Nijar)|ne',
  },
  {
    label: 'Nigeria',
    value: 'Nigeria|ng',
  },
  {
    label: 'Niue',
    value: 'Niue|nu',
  },
  {
    label: 'Norfolk Island',
    value: 'Norfolk Island|nf',
  },
  {
    label: 'North Korea (조선 민주주의 인민 공화국)',
    value: 'North Korea (조선 민주주의 인민 공화국)|kp',
  },
  {
    label: 'Northern Mariana Islands',
    value: 'Northern Mariana Islands|mp',
  },
  {
    label: 'Norway (Norge)',
    value: 'Norway (Norge)|no',
  },
  {
    label: 'Oman (عُمان‎)',
    value: 'Oman (عُمان‎)|om',
  },
  {
    label: 'Pakistan (پاکستان‎)',
    value: 'Pakistan (پاکستان‎)|pk',
  },
  {
    label: 'Palau',
    value: 'Palau|pw',
  },
  {
    label: 'Palestine (فلسطين‎)',
    value: 'Palestine (فلسطين‎)|ps',
  },
  {
    label: 'Panama (Panamá)',
    value: 'Panama (Panamá)|pa',
  },
  {
    label: 'Papua New Guinea',
    value: 'Papua New Guinea|pg',
  },
  {
    label: 'Paraguay',
    value: 'Paraguay|py',
  },
  {
    label: 'Peru (Perú)',
    value: 'Peru (Perú)|pe',
  },
  {
    label: 'Philippines',
    value: 'Philippines|ph',
  },
  {
    label: 'Poland (Polska)',
    value: 'Poland (Polska)|pl',
  },
  {
    label: 'Portugal',
    value: 'Portugal|pt',
  },
  {
    label: 'Puerto Rico',
    value: 'Puerto Rico|pr',
  },
  {
    label: 'Qatar (قطر‎)',
    value: 'Qatar (قطر‎)|qa',
  },
  {
    label: 'Réunion (La Réunion)',
    value: 'Réunion (La Réunion)|re',
  },
  {
    label: 'Romania (România)',
    value: 'Romania (România)|ro',
  },
  {
    label: 'Russia (Россия)',
    value: 'Russia (Россия)|ru',
  },
  {
    label: 'Rwanda',
    value: 'Rwanda|rw',
  },
  {
    label: 'Saint Barthélemy',
    value: 'Saint Barthélemy|bl',
  },
  {
    label: 'Saint Helena',
    value: 'Saint Helena|sh',
  },
  {
    label: 'Saint Kitts and Nevis',
    value: 'Saint Kitts and Nevis|kn',
  },
  {
    label: 'Saint Lucia',
    value: 'Saint Lucia|lc',
  },
  {
    label: 'Saint Martin (Saint-Martin (partie française))',
    value: 'Saint Martin (Saint-Martin (partie française))|mf',
  },
  {
    label: 'Saint Pierre and Miquelon (Saint-Pierre-et-Miquelon)',
    value: 'Saint Pierre and Miquelon (Saint-Pierre-et-Miquelon)|pm',
  },
  {
    label: 'Saint Vincent and the Grenadines',
    value: 'Saint Vincent and the Grenadines|vc',
  },
  {
    label: 'Samoa',
    value: 'Samoa|ws',
  },
  {
    label: 'San Marino',
    value: 'San Marino|sm',
  },
  {
    label: 'São Tomé and Príncipe (São Tomé e Príncipe)',
    value: 'São Tomé and Príncipe (São Tomé e Príncipe)|st',
  },
  {
    label: 'Saudi Arabia (المملكة العربية السعودية‎)',
    value: 'Saudi Arabia (المملكة العربية السعودية‎)|sa',
  },
  {
    label: 'Senegal (Sénégal)',
    value: 'Senegal (Sénégal)|sn',
  },
  {
    label: 'Serbia (Србија)',
    value: 'Serbia (Србија)|rs',
  },
  {
    label: 'Seychelles',
    value: 'Seychelles|sc',
  },
  {
    label: 'Sierra Leone',
    value: 'Sierra Leone|sl',
  },
  {
    label: 'Singapore',
    value: 'Singapore|sg',
  },
  {
    label: 'Sint Maarten',
    value: 'Sint Maarten|sx',
  },
  {
    label: 'Slovakia (Slovensko)',
    value: 'Slovakia (Slovensko)|sk',
  },
  {
    label: 'Slovenia (Slovenija)',
    value: 'Slovenia (Slovenija)|si',
  },
  {
    label: 'Solomon Islands',
    value: 'Solomon Islands|sb',
  },
  {
    label: 'Somalia (Soomaaliya)',
    value: 'Somalia (Soomaaliya)|so',
  },
  {
    label: 'South Africa',
    value: 'South Africa|za',
  },
  {
    label: 'South Korea (대한민국)',
    value: 'South Korea (대한민국)|kr',
  },
  {
    label: 'South Sudan (جنوب السودان‎)',
    value: 'South Sudan (جنوب السودان‎)|ss',
  },
  {
    label: 'Spain (España)',
    value: 'Spain (España)|es',
  },
  {
    label: 'Sri Lanka (ශ්‍රී ලංකාව)',
    value: 'Sri Lanka (ශ්‍රී ලංකාව)|lk',
  },
  {
    label: 'Sudan (السودان‎)',
    value: 'Sudan (السودان‎)|sd',
  },
  {
    label: 'Suriname',
    value: 'Suriname|sr',
  },
  {
    label: 'Svalbard and Jan Mayen',
    value: 'Svalbard and Jan Mayen|sj',
  },
  {
    label: 'Swaziland',
    value: 'Swaziland|sz',
  },
  {
    label: 'Sweden (Sverige)',
    value: 'Sweden (Sverige)|se',
  },
  {
    label: 'Switzerland (Schweiz)',
    value: 'Switzerland (Schweiz)|ch',
  },
  {
    label: 'Syria (سوريا‎)',
    value: 'Syria (سوريا‎)|sy',
  },
  {
    label: 'Taiwan (台灣)',
    value: 'Taiwan (台灣)|tw',
  },
  {
    label: 'Tajikistan',
    value: 'Tajikistan|tj',
  },
  {
    label: 'Tanzania',
    value: 'Tanzania|tz',
  },
  {
    label: 'Thailand (ไทย)',
    value: 'Thailand (ไทย)|th',
  },
  {
    label: 'Timor-Leste',
    value: 'Timor-Leste|tl',
  },
  {
    label: 'Togo',
    value: 'Togo|tg',
  },
  {
    label: 'Tokelau',
    value: 'Tokelau|tk',
  },
  {
    label: 'Tonga',
    value: 'Tonga|to',
  },
  {
    label: 'Trinidad and Tobago',
    value: 'Trinidad and Tobago|tt',
  },
  {
    label: 'Tunisia (تونس‎)',
    value: 'Tunisia (تونس‎)|tn',
  },
  {
    label: 'Turkey (Türkiye)',
    value: 'Turkey (Türkiye)|tr',
  },
  {
    label: 'Turkmenistan',
    value: 'Turkmenistan|tm',
  },
  {
    label: 'Turks and Caicos Islands',
    value: 'Turks and Caicos Islands|tc',
  },
  {
    label: 'Tuvalu',
    value: 'Tuvalu|tv',
  },
  {
    label: 'U.S. Virgin Islands',
    value: 'U.S. Virgin Islands|vi',
  },
  {
    label: 'Uganda',
    value: 'Uganda|ug',
  },
  {
    label: 'Ukraine (Україна)',
    value: 'Ukraine (Україна)|ua',
  },
  {
    label: 'United Arab Emirates (الإمارات العربية المتحدة‎)',
    value: 'United Arab Emirates (الإمارات العربية المتحدة‎)|ae',
  },
  {
    label: 'United Kingdom',
    value: 'United Kingdom|gb',
  },
  {
    label: 'United States',
    value: 'United States|us',
  },
  {
    label: 'Uruguay',
    value: 'Uruguay|uy',
  },
  {
    label: 'Uzbekistan (Oʻzbekiston)',
    value: 'Uzbekistan (Oʻzbekiston)|uz',
  },
  {
    label: 'Vanuatu',
    value: 'Vanuatu|vu',
  },
  {
    label: 'Vatican City (Città del Vaticano)',
    value: 'Vatican City (Città del Vaticano)|va',
  },
  {
    label: 'Venezuela',
    value: 'Venezuela|ve',
  },
  {
    label: 'Vietnam (Việt Nam)',
    value: 'Vietnam (Việt Nam)|vn',
  },
  {
    label: 'Wallis and Futuna (Wallis-et-Futuna)',
    value: 'Wallis and Futuna (Wallis-et-Futuna)|wf',
  },
  {
    label: 'Western Sahara (الصحراء الغربية‎)',
    value: 'Western Sahara (الصحراء الغربية‎)|eh',
  },
  {
    label: 'Yemen (اليمن‎)',
    value: 'Yemen (اليمن‎)|ye',
  },
  {
    label: 'Zambia',
    value: 'Zambia|zm',
  },
  {
    label: 'Zimbabwe',
    value: 'Zimbabwe|zw',
  },
  {
    label: 'Åland Islands',
    value: 'Åland Islands|ax',
  },
];

const executionClientData = [
  { label: 'Geth', value: 'Geth' },
  { label: 'Nethermind', value: 'Nethermind' },
  { label: 'Besu', value: 'Besu' },
  { label: 'Erigon', value: 'Erigon' },
  { label: 'Akula', value: 'Akula' },
];

const consensusClientData = [
  { label: 'Lighthouse', value: 'Lighthouse' },
  { label: 'Teku', value: 'Teku' },
  { label: 'Nimbus', value: 'Nimbus' },
  { label: 'Lodestar', value: 'Lodestar' },
  { label: 'Prysm', value: 'Prysm' },
];

const cloudProviderData = [
  { label: 'Amazon Web Services (AWS)', value: 'Amazon Web Services (AWS)' },
  { label: 'Microsoft Azure', value: 'Microsoft Azure' },
  {
    label: 'Google Cloud Platform (GCP)',
    value: 'Google Cloud Platform (GCP)',
  },
  { label: 'Alibaba Cloud', value: 'Alibaba Cloud' },
  { label: 'Oracle Cloud', value: 'Oracle Cloud' },
  { label: 'IBM Cloud (Kyndryl)', value: 'IBM Cloud (Kyndryl)' },
  { label: 'Tencent Cloud', value: 'Tencent Cloud' },
  { label: 'OVHcloud', value: 'OVHcloud' },
  { label: 'DigitalOcean', value: 'DigitalOcean' },
  { label: 'Linode (Akamai)', value: 'Linode (Akamai)' },
];

export {
  relaysSupportedData,
  locationData,
  executionClientData,
  consensusClientData,
  cloudProviderData,
};
