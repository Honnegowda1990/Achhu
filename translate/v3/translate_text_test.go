// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestTranslateText(t *testing.T) {
	t.Skip("Skipped while investigating https://github.com/GoogleCloudPlatform/golang-samples/issues/2811")
	tc := testutil.SystemTest(t)

	sourceLang := "en-US"
	targetLang := "sr-Latn"
	text := "Hello world"

	// Translate text.
	var buf bytes.Buffer
	if err := translateText(&buf, tc.ProjectID, sourceLang, targetLang, text); err != nil {
		t.Fatalf("translateText: %v", err)
	}
	if got, want1, want2 := buf.String(), "Zdravo", "Pozdrav"; !strings.Contains(got, want1) && !strings.Contains(got, want2) {
		t.Errorf("translateText got:\n----\n%s----\nWant to contain:\n----\n%s\n----\nOR\n----\n%s\n----", got, want1, want2)
	}
}
