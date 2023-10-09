import 'package:flutter_test/flutter_test.dart';
import 'package:crypto_export/crypto_export.dart';

void main() {
  var crypt = Crypt();
  test('getKeysTest', () {
    var keys = crypt.keys();
    expect(keys.length, 4);
  });

  test('getKeysIos', () {
    
  });
}
