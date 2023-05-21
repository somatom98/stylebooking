import 'package:flutter/material.dart';
import 'package:stylebooking_app/src/models/product.dart';

// ignore: must_be_immutable
class Products extends StatelessWidget {
  List<ProductViewModel> products;

  Products(this.products, {super.key});

  @override 
  Widget build(BuildContext context){
    return SingleChildScrollView(
      child: Column(
        children:products.map(
          (x) => GestureDetector(
            onTap: () {
              
            },
            child: Card(
              child: Padding (
                padding: const EdgeInsets.all(32),
                child: Stack(
                  alignment: Alignment.center,
                  children: <Widget>[
                    Image.asset('assets/${x.name.toLowerCase()}.png'),
                    Text(
                      x.name,
                      style: const TextStyle(
                        fontSize: 64,
                        fontWeight: FontWeight.bold,
                        color: Colors.white,
                      ),
                    )
                  ]
                ),
              )
            )
          )
        ).toList()
      )
    );
  }
}