CREATE ROLE cinemavault WITH LOGIN PASSWORD 'pa55word';
GRANT ALL PRIVILEGES ON DATABASE cinemavault TO cinemavault;
\c cinemavault postgres;
CREATE EXTENSION IF NOT EXISTS citext;
GRANT ALL ON SCHEMA public TO cinemavault;
