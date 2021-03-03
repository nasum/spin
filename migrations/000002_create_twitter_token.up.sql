CREATE TABLE IF NOT EXISTS twitter_tokens(
  id serial PRIMARY KEY,
  user_id int NOT NULL,
  access_token text NOT NULL,
  access_token_seacret text NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
