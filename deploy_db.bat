@echo off
pg_dump -U postgres -h localhost -c quillard > quillard.sql
psql -U postgres -h platinium.ddns.net quillard < quillard.sql