OUT=jp

run: clean build
	# Check if out is not empty.
	@if [ -z "${out}" ]; then \
		echo "out should be empty; ex: make out=ddMMyyyy"; \
		exit 1; \
	fi
	# Check if Mf_portfolio.json exists.
	@if [ -z "~/franklin_json/Mf_portfolio.json" ]; then \
		echo "copy Mf_portfolio to from '~/franklin_json' first."; \
		exit 1; \
	fi

	# Create directory for Mf_portfolio's full and franklin fund.
	mkdir -p ~/franklin_json/full_response
	mkdir -p ~/franklin_json/franklin

	# Extracting franklin's data
	./${OUT}

	# Formatting & storing franklin's data.
	touch ~/franklin_json/franklin/${out}_franklin.json
	jq . ~/franklin_json/outFile.json > ~/franklin_json/franklin/${out}_franklin.json

	# Removing temp file.
	rm ~/franklin_json/outFile.json
	
	# Format and Move full response.
	jq . ~/franklin_json/Mf_portfolio.json > ~/franklin_json/${out}.json
	mv ~/franklin_json/${out}.json ~/franklin_json/full_response/${out}_full_response.json

	# Remove the unformatted full response.
	rm ~/franklin_json/Mf_portfolio.json

clean:
	rm ./${OUT}

build: *.go
	go build -o ${OUT} .



