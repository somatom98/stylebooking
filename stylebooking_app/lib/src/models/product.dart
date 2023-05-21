import 'dart:ffi';

class ProductViewModel {
  final String name;
  final String description;
  final String price;

  ProductViewModel({
    required this.name,
    required this.description,
    required this.price,
  });

  factory ProductViewModel.fromJson(Map<String, dynamic> json) {
    return ProductViewModel(
      name: json['name'],
      description: json['description'],
      price: json['price'],
    );
  }
}

class ProductListViewModel {
  final List<ProductViewModel> products;

  ProductListViewModel({
    required this.products,
  });

  factory ProductListViewModel.fromJson(List<dynamic> json) {
    return ProductListViewModel(
      products: json.map((x) => ProductViewModel.fromJson(x)).toList(),
    );
  }
}