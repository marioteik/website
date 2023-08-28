DATABASE_URL=postgresql://postgres:mqiC6j2YuEboCVp1@db.wlqronusgmkqccajzbqt.supabase.co:5432/postgres
SUPABASE_URL=https://wlqronusgmkqccajzbqt.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndscXJvbnVzZ21rcWNjYWp6YnF0Iiwi>
API_PORT=4001
BINARIES=dist

# Serve dev environment
dev: start.tailwind start.website start.browserSync

start.website:
	@echo "Starting API server..."
	@air -c .air.toml

start.tailwind:
	@echo "Starting Tailwind server..."
	@npm run tailwind:watch

start.browserSync:
	@echo "Starting BrowserSync..."
	@browser-sync start --proxy "localhost:4000" --files "public/**/*" --files "components/**/*.gohtml" --files "pages/**/*.gohtml" --reloadDebounce 500 --browser "google chrome" --port 4100