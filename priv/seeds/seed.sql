INSERT INTO snippets (title, content, expires) VALUES (
  'An old silent pond',
  'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
  current_date + interval '365 days');

INSERT INTO snippets (title, content,  expires) VALUES 
('Over the wintry forest',
  'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
  current_date + interval '7 days'); 

INSERT INTO snippets (title, content,  expires) VALUES (
  'First autumn morning',
  'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
   current_date + interval '7 days');
