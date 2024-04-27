// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/version"
)

var NotFoundPagePath = ""

const (
	NotFound = `<!-- @format -->

	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta
		  name="viewport"
		  content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"
		/>
		<title>GaiaNet</title>
	  </head>
	  <body>
		<div
		  style="
			width: 100%;
			height: 100vh;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			background: #ffffff;
			font-family: sans-serif;
		  "
		>
		  <svg
			width="222"
			height="271"
			viewBox="0 0 222 271"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
		  >
			<path d="M35 238L1.5 269H211L222 261.5V238H35Z" fill="#D43327" />
			<path d="M2 118.5L8.5 107L7.5 263.5L2 269V118.5Z" fill="#902552" />
			<path
			  d="M220 106H221L221 107L221 263L221 263.551L220.534 263.845L211.034 269.845L210.789 270H210.5H1.99982H0.999817V269V120.542V120.314L1.0983 120.109L7.5983 106.567L7.87057 106H8.49982H220Z"
			  stroke="black"
			  stroke-width="2"
			/>
			<rect x="8.25" y="106.75" width="212.5" height="156" fill="#F7F7F7" />
			<rect
			  x="8.25"
			  y="106.75"
			  width="212.5"
			  height="156"
			  stroke="black"
			  stroke-width="1.5"
			/>
			<path
			  d="M9.49882 107.5H220L219.501 262H9L9.49882 107.5Z"
			  fill="white"
			/>
			<line
			  x1="17"
			  y1="117.25"
			  x2="27"
			  y2="117.25"
			  stroke="black"
			  stroke-width="1.5"
			/>
			<line
			  x1="17"
			  y1="121.25"
			  x2="27"
			  y2="121.25"
			  stroke="black"
			  stroke-width="1.5"
			/>
			<line
			  x1="17"
			  y1="125.25"
			  x2="27"
			  y2="125.25"
			  stroke="black"
			  stroke-width="1.5"
			/>
			<circle cx="198.5" cy="121.5" r="3" stroke="black" />
			<circle cx="208.5" cy="121.5" r="3" stroke="black" />
			<circle cx="174" cy="117" r="3.5" stroke="black" />
			<circle cx="55" cy="117" r="3.5" stroke="black" />
			<g opacity="0.3">
			  <path
				d="M146.984 120.609L144.096 136.352L171.129 159.168L174.005 143.465L146.984 120.609Z"
				fill="#D43327"
			  />
			  <path
				d="M98.6487 152.635L126.383 176.065L117.516 225.783L89.9543 202.442L98.6487 152.635Z"
				fill="#D43327"
			  />
			  <path
				d="M56.8483 227.52L84.0407 250.677L159.979 223.021L132.763 200.035L56.8483 227.52Z"
				fill="#902552"
			  />
			  <path
				d="M98.6621 152.648L126.075 175.729L171.127 159.168L144.097 136.339L98.6621 152.648Z"
				fill="#902552"
			  />
			  <path
				d="M141.361 152.304L168.366 175.111L159.946 223.057L132.729 200.071L141.361 152.304Z"
				fill="#D43327"
			  />
			  <path
				d="M146.987 120.586L144.1 136.344L98.6522 152.635L90.8776 196.926L106.24 191.428L111.282 163.316L141.378 152.284L132.76 200.034L56.7733 227.521L70.8712 148.049L146.987 120.586Z"
				fill="#F3D8B9"
			  />
			  <path
				d="M106.279 191.382L111.313 163.271L141.396 152.263L132.764 200.03L56.8126 227.483L70.9063 148.008L147.02 120.564L144.131 136.299L98.6998 152.604L90.9128 196.884L106.279 191.382Z"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			  <path
				d="M98.7004 152.603L111.314 163.27"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			  <path
				d="M141.731 151.866L168.521 174.491L159.655 223.401L83.787 250.974L56.5293 227.819"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			  <path
				d="M132.766 200.03L159.982 223.016"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			  <path
				d="M147.388 120.24L173.909 142.639L171.049 159.263L156.649 164.461"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			  <path
				d="M144.133 136.299L171.163 159.127"
				stroke="#020000"
				stroke-width="1.08937"
				stroke-miterlimit="10"
			  />
			</g>
			<g clip-path="url(#clip0_357_688)">
			  <path
				d="M86.22 197.772H77.7785V206.214H69.337V197.772H44V180.883H52.4415V189.325H69.3308V172.442H60.8892V164H77.7785V189.331H86.22V197.772ZM52.4415 172.442H60.883V180.883H52.4415V172.442Z"
				fill="black"
			  />
			  <path
				d="M94.6621 197.772V172.442H103.104V197.772H94.6621ZM103.104 164H128.434V172.442H103.104V164ZM128.434 206.22H103.104V197.778H128.434V206.22ZM136.876 197.778H128.434V172.442H136.876V197.772V197.778Z"
				fill="black"
			  />
			  <path
				d="M187.544 197.772H179.102V206.214H170.661V197.772H145.33V180.883H153.772V189.325H170.661V172.442H162.219V164H179.109V189.331H187.55V197.772H187.544ZM153.772 172.442H162.213V180.883H153.772V172.442Z"
				fill="black"
			  />
			</g>
			<path
			  d="M54.5 117L110 17.5"
			  stroke="black"
			  stroke-width="1.5"
			  stroke-linecap="round"
			/>
			<path
			  d="M174.5 117L119 17.5"
			  stroke="black"
			  stroke-width="1.5"
			  stroke-linecap="round"
			/>
			<mask
			  id="path-30-outside-1_357_688"
			  maskUnits="userSpaceOnUse"
			  x="108"
			  y="6"
			  width="13"
			  height="21"
			  fill="black"
			>
			  <rect fill="white" x="108" y="6" width="13" height="21" />
			  <path
				fill-rule="evenodd"
				clip-rule="evenodd"
				d="M119 8H110V20.5V21H110.027C110.276 23.25 112.184 25 114.5 25C116.816 25 118.724 23.25 118.973 21H119V20.5V8Z"
			  />
			</mask>
			<path
			  fill-rule="evenodd"
			  clip-rule="evenodd"
			  d="M119 8H110V20.5V21H110.027C110.276 23.25 112.184 25 114.5 25C116.816 25 118.724 23.25 118.973 21H119V20.5V8Z"
			  fill="#D9D9D9"
			/>
			<path
			  d="M110 8V6.5H108.5V8H110ZM119 8H120.5V6.5H119V8ZM110 21H108.5V22.5H110V21ZM110.027 21L111.518 20.8352L111.371 19.5H110.027V21ZM118.973 21V19.5H117.629L117.482 20.8352L118.973 21ZM119 21V22.5H120.5V21H119ZM110 9.5H119V6.5H110V9.5ZM111.5 20.5V8H108.5V20.5H111.5ZM111.5 21V20.5H108.5V21H111.5ZM110.027 19.5H110V22.5H110.027V19.5ZM114.5 23.5C112.957 23.5 111.684 22.3336 111.518 20.8352L108.537 21.1648C108.868 24.1664 111.411 26.5 114.5 26.5V23.5ZM117.482 20.8352C117.316 22.3336 116.043 23.5 114.5 23.5V26.5C117.589 26.5 120.132 24.1664 120.463 21.1648L117.482 20.8352ZM119 19.5H118.973V22.5H119V19.5ZM117.5 20.5V21H120.5V20.5H117.5ZM117.5 8V20.5H120.5V8H117.5Z"
			  fill="black"
			  mask="url(#path-30-outside-1_357_688)"
			/>
			<circle
			  cx="114.5"
			  cy="6.5"
			  r="5.25"
			  fill="white"
			  stroke="black"
			  stroke-width="1.5"
			/>
			<defs>
			  <clipPath id="clip0_357_688">
				<rect
				  width="143.543"
				  height="42.22"
				  fill="white"
				  transform="translate(44 164)"
				/>
			  </clipPath>
			</defs>
		  </svg>
	
		  <h2
			style="
			  color: #000;
			  font-size: 40px;
			  font-weight: 700;
			  margin: 40px 0 0 0;
			  text-align: center;
			  padding: 0 16px;
			"
		  >
			GaiaNet node is not started
		  </h2>
	
		  <p
			style="
			  margin: 40px 0 0 0;
			  color: #555;
			  font-size: 14px;
			  font-weight: 400;
			  text-align: center;
			  padding: 0 16px;
			"
		  >
			Visit
			<a href="https://www.gaianet.ai" target="_blank" style="color: #d43327"
			  >https://www.gaianet.ai</a
			>
			for more information
		  </p>
		</div>
	  </body>
	</html>	
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = os.ReadFile(NotFoundPagePath)
		if err != nil {
			log.Warnf("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func NotFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	content := getNotFoundPageContent()
	res := &http.Response{
		Status:        "Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	return res
}
