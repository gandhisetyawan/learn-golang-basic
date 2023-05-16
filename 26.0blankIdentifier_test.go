/** Blank Identifier
- Kadang kita hanya perlu menjalankan init function pada package
	tanpa harus menjalankan function yang ada pada package
- secara default, golang akan komplain ketika ada package yang di import
	namun tidak di gunakan
*/

package learn_golang_basic

import (
	_ "learn_golang_basic/database"
	"testing"
)

func TestBlankIdentifier(t *testing.T) {
	// database.GetDatabase()

}
