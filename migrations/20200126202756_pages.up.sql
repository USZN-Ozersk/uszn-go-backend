CREATE TABLE pages (
    page_id SERIAL NOT NULL PRIMARY KEY,
    page_name VARCHAR NOT NULL,
    page_text TEXT NOT NULL,
    page_menu INTEGER UNIQUE REFERENCES menu (menu_id)
);

INSERT INTO pages (page_name, page_text, page_menu)
VALUES
(
    'Страница 1',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    1
),
(
    'Страница 2',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    2
),
(
    'Страница 3',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    3
),
(
    'Страница 4',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    4
),
(
    'Страница 5',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    5
),
(
    'Страница 6',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    6
),
(
    'Страница 7',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    7
),
(
    'Страница 8',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    8
),
(
    'Страница 9',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    9
),
(
    'Страница 10',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    10
),
(
    'Страница 11',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    11
),
(
    'Страница 12',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    12
);