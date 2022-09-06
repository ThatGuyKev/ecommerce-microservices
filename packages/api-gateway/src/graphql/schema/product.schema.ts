import { gql } from "apollo-server-core";

export default gql`
  interface Node {
    id: ID!
  }
  input PaginationInput {
    before: String
    after: String
    first: Float
    last: Float
  }
  type PageInfo {
    hasNextPage: Boolean!
    hasPreviousPage: Boolean!
    startCursor: String
    endCursor: String
  }
  type PageData {
    count: Float
    limit: Float
    ofset: Float
  }
  type ProductBasic implements Node {
    id: ID!
    title: String
    defaultPrice: PriceInfo
    currency: String
    reviewsSummary: ReviewsSummary
    tags: [String]
    primaryAsset: Asset
  }
  type ReviewsSummary {
    numberOfReviews: Int
    rating: Float
    ratingUnits: String
  }
  type Asset {
    isPrimary: Boolean
    sorting: String
    type: String
    url: String
    altText: String
    title: String
  }

  type Product {
    id: ID!
    title: String
    uri: String
    description: String
    shortDescription: String
    attributes: [Attribute]
    sku: String
    upc: String
    currency: String
    defaultPrice: PriceInfo
    salePrice: PriceInfo
    metaTitle: String
    metaDescription: String
    displayTemplate: String
    keywords: [String]
    discountable: Boolean
    reviewsSummary: ReviewsSummary
    options: [ProductOption]
    primaryAsset: Asset
    PrimaryCategory: Category
    brand: String
  }

  type PriceInfo {
    currency: String
    amount: Float
  }

  type ProductOption {
    id: String
    label: String
    type: String
    displayOrder: Int
    attributeChoice: AttributeChoice
    templateContextId: String
    itemChoice: ItemChoice
  }
  type ItemChoice {
    targetType: String
    selectionType: String # enum
    choiceKey: String
    maximumQuantity: Int
    minimumQuantity: Int
    pricingModel: String
    overridePrice: PriceInfo
    discountAllowed: Boolean
    specificChoices: [SpecificChoice]
  }
  type Attribute {
    key: String
    value: String
  }
  type SpecificChoice {
    type: String
    nameOverride: String
    pricingModel: String
    overridePrice: PriceInfo
    discountAllowed: Boolean
    asset: Asset
  }

  type AttributeChoice {
    attributeName: String
    type: String
    required: Boolean
    errorMessage: String
    errorCode: String
    allowedValues: [OptionValue]
  }
  type OptionValue {
    id: String
    label: String
    value: String
    displayOrder: Int
  }
  type ProductResponse {
    page: ProductConnection!
    pageData: PageData
  }
  type ProductConnection {
    edges: [ProductEdge!]
    pageInfo: PageInfo!
  }
  type ProductEdge {
    cursor: String!
    node: Product
  }

  type CategoryBasic {
    id: ID!
    title: String
    url: String
  }
  type Category {
    id: ID!
    title: String
    url: String
    description: String
    displayTemplate: String
    metaTitle: String
    metaDescription: String
    attributes: [Attribute]
    parentCategory: CategoryBasic
  }
  type Query {
    category(id: String!): Category
    categories(ids: [String]): [Category]
    subcategories(categoryId: String!): [Category]
    categoryAssets(categoryId: String!): [Asset]
    product(productId: String!): Product
    products(ids: [String], pagination: PaginationInput): [ProductBasic]
    productAssets(productId: String!): [Asset]
    categoryProducts(
      categoryId: String!
      pagination: PaginationInput
    ): [ProductBasic]
  }
`;
