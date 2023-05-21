import 'dart:ffi';

import 'package:stylebooking_app/product_manager.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'bottom_bar.dart';

void main() {
  runApp(ProviderScope(child: MyApp()));
}

final pageProvider = StateNotifierProvider<PageNotifier, int>((ref) {
  return PageNotifier(page: 1);
});

class PageNotifier extends StateNotifier<int> {
  PageNotifier({page}) : super(page);

  void selectPage(int page) {
    state = page;
  }
}

class MyApp extends ConsumerWidget{
  const MyApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var page = ref.watch(pageProvider);

    return MaterialApp(
      theme: ThemeData(
        // brightness: Brightness.dark,
        // primaryColor: const Color(0xFFBE3455),
        // indicatorColor: const Color(0xFF34BF49),
        // disabledColor: const Color(0xFF087318),
        // splashColor: const Color(0xFFFF5E84),
        // highlightColor: const Color(0xFF73142A), 
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFFBE3455),
          brightness: Brightness.light,
          secondary: const Color(0xFF34BF49)
        ),
      ),
      home: Scaffold(
        body: _loadPage(page),
        bottomNavigationBar: const BottomBar(),
      ),
    );
  }

  Widget _loadPage(int page) {
    switch (page) {
      case 0:
        return const ProductManager();
      case 1: 
        return const ProductManager();
      case 2:
        return const ProductManager();
      default:
        return Text('$page');
    }
  }
}