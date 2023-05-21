import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'main.dart';

class BottomBar extends ConsumerWidget {
  const BottomBar({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return BottomNavigationBar(
      showUnselectedLabels: false,
      currentIndex: ref.watch(pageProvider),
      items: const [
        BottomNavigationBarItem(
          icon: Icon(Icons.search), 
          label: 'Search',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.favorite_rounded), 
          label: 'Favourites'
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.person_2_rounded), 
          label: 'Profile'
        ),
      ],
      onTap: (value) {
        ref.read(pageProvider.notifier).selectPage(value);
      },
    );
  }
}