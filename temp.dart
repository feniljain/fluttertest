import 'package:flutter_test/flutter_test.dart';

import '../lib/fibo.dart';

void main() {
  group('Test Get1', () {
    test('Fibo Number At Pos', () {
      //arrange
      var fibo = Fibo();
      //act
      var result = fibo.getFiboNumberAtPost(5);
      //assert
      expect(result, 8);
    });
  });}