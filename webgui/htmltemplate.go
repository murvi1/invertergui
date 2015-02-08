/*
Copyright (c) 2015, Hendrik van Wyk
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

* Neither the name of invertergui nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package webgui

var htmlTemplate string = `<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="refresh" content="5">

    <title>Victron Multiplus Monitor</title>
</head>

<body>
	{{if .Error}} <p>Error encountered: {{.Error}} </p> {{end}}
	<dl>
		<dt> LEDs:</dt>
		{{range .Leds}}
		<dt> {{.}}</dt>
		{{else}}
		None
		{{end}}
	</dl>
	<dl>
		<dt> Output Current: {{.OutCurrent}} A</dt>
		<dt> Output Voltage: {{.OutVoltage}} V</dt>
		<dt> Output Power: {{.OutPower}} VA</dt>
	</dl>
	<dl>
		<dt> Input Current: {{.InCurrent}} A</dt>
		<dt> Input Voltage: {{.InVoltage}} V</dt>
		<dt> Input Frequency: {{.InFreq}} Hz</dt>
		<dt> Input Power: {{.InPower}} VA</dt>
	</dl>
	<p> Input - Output Power: {{.InMinOut}} VA</p>
	<dl>
		<dt> Battery Current: {{.BatCurrent}} A</dt>
		<dt> Battery Voltage: {{.BatVoltage}} V</dt>
		<dt> Battery Power: {{.BatPower}} W</dt>
	</dl>
</body>
</html>`
