CREATE TABLE news (
    news_id SERIAL PRIMARY KEY NOT NULL,
    news_name VARCHAR NOT NULL,
    news_date DATE NOT NULL DEFAULT current_date,
    news_text TEXT NOT NULL,
    news_img TEXT NOT NULL
);

INSERT INTO news (news_name, news_text, news_img)
VALUES
(
    'Новость 1',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_01.jpg'
),
(
    'Новость 1',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_04.jpg'
),
(
    'Новость 2',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_02.jpg'
),
(
    'Новость 3',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_03.jpg'
),
(
    'Новость 4',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_01.jpg'
),
(
    'Новость 5',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_04.jpg'
),
(
    'Новость 6',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_01.jpg'
),
(
    'Новость 7',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_02.jpg'
),
(
    'Новость 8',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_03.jpg'
),
(
    'Новость 9',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_04.jpg'
),
(
    'Новость 10',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_01.jpg'
),
(
    'Новость 11',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_02.jpg'
),
(
    'Новость 12',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_03.jpg'
),
(
    'Новость 13',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque lobortis lacus laoreet, convallis diam eu, aliquam purus. Fusce mollis feugiat lacus id commodo. Etiam libero purus, interdum eget dolor nec, laoreet suscipit ex. Sed accumsan, dolor sit amet scelerisque cursus, orci arcu semper leo, quis molestie ligula ex non erat. In pharetra rutrum nulla semper tincidunt. In efficitur condimentum odio ut tincidunt. Maecenas auctor, dui sed sagittis sagittis, urna eros facilisis leo, a dictum nisi purus ac erat. Quisque dignissim elit ac velit suscipit, ut tincidunt quam cursus. Cras tincidunt maximus sapien, eu sagittis felis tempus sed. Donec lorem augue, laoreet sit amet justo nec, dapibus aliquet neque. Vivamus commodo commodo risus, nec vulputate velit dictum id. Fusce ultricies consequat nisl sagittis luctus. Nullam blandit mauris a eleifend viverra.',
    '20200211/img_04.jpg'
);