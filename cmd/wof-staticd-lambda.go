package main

import (
	"github.com/whosonfirst/algnhsa"
	"github.com/whosonfirst/go-whosonfirst-readwrite-s3/config"
	"github.com/whosonfirst/go-whosonfirst-readwrite-s3/reader"
	"github.com/whosonfirst/go-whosonfirst-static/http"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	// PLEASE FOR TO BE READING FROM OTHER SOURCES (GITHUB, HTTP, FS)

	// https://docs.aws.amazon.com/lambda/latest/dg/env_variables.html
	// WHOSONFIRST_STATIC_S3_DSN="bucket={BUCKET} prefix={PREFIX} region={REGION} credentials={CREDS}"

	s3_dsn := os.Getenv("WHOSONFIRST_STATIC_S3_DSN")

	cfg, err := config.NewS3ConfigFromString(s3_dsn)

	if err != nil {
		log.Fatal(err)
	}

	r, err := reader.NewS3Reader(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// PLEASE FOR TO BE SETTING THESE FROM ENVIRONMENT VARIABLES

	enable_all := false
	enable_html := false
	enable_graphics := false
	enable_data := false
	enable_png := true
	enable_svg := true
	enable_spr := true
	enable_geojson := true

	var png_handler gohttp.Handler
	var svg_handler gohttp.Handler

	var geojson_handler gohttp.Handler
	var spr_handler gohttp.Handler

	binary_types := make([]string, 0)

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	if enable_all || enable_html {
		// please add me...
	}

	if enable_all || enable_graphics || enable_png {

		png_opts, err := http.NewDefaultRasterOptions()

		if err != nil {
			log.Fatal(err)
		}

		h, err := http.RasterHandler(r, png_opts)

		if err != nil {
			log.Fatal(err)
		}

		png_handler = h
		mux.Handle("/png/", png_handler)

		binary_types = append(binary_types, "image/png")
	}

	if enable_all || enable_graphics || enable_svg {

		svg_opts, err := http.NewDefaultSVGOptions()

		if err != nil {
			log.Fatal(err)
		}

		h, err := http.SVGHandler(r, svg_opts)

		if err != nil {
			log.Fatal(err)
		}

		svg_handler = h
		mux.Handle("/svg/", svg_handler)
	}

	if enable_all || enable_data || enable_spr {

		h, err := http.SPRHandler(r)

		if err != nil {
			log.Fatal(err)
		}

		spr_handler = h
		mux.Handle("/spr/", spr_handler)
	}

	if enable_all || enable_data || enable_geojson {

		h, err := http.GeoJSONHandler(r)

		if err != nil {
			log.Fatal(err)
		}

		geojson_handler = h
		mux.Handle("/geojson/", geojson_handler)
	}

	if enable_all || enable_html {
		// please write me
	}

	opts := new(algnhsa.Options)
	opts.BinaryContentTypes = binary_types

	algnhsa.ListenAndServe(mux, opts)
}
