INSERT INTO apps(id, name, secret)
VALUE(1, 'test', 'test-secret')
ON CONFLICT DO NOTHING;
