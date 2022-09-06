export default {
  id: "string",
  title: "string",
  uri: "string",
  description: "string",
  shortDescription: "string",
  sku: "string",
  upc: "string",
  currency: "USD",
  defaultPrice: {
    amount: 0,
    currency: "USD",
  },
  salePrice: {
    amount: 0,
    currency: "USD",
  },
  metaTitle: "string",
  metaDescription: "string",
  displayTemplate: "string",
  keywords: ["string"],
  discountable: true,
  reviewsSummary: {
    numberOfReviews: 0,
    rating: 0,
    ratingUnits: "string",
  },
  options: [
    {
      id: "string",
      label: "string",
      type: "VARIANT_DISTINGUISHING",
      displayOrder: 0,
      attributeChoice: {
        attributeName: "string",
        type: "COLOR",
        required: false,

        errorMessage: "string",
        errorCode: "string",
        allowedValues: [
          {
            id: "string",
            label: "string",
            value: "string",
            displayOrder: 0,
          },
        ],
      },
      templateContextId: "string",
      itemChoice: {
        targetType: "SPECIFIC_PRODUCTS",
        selectionType: "CHOOSE_ONE",
        choiceKey: "string",
        maximumQuantity: 0,
        minimumQuantity: 0,
        pricingModel: "string",
        overridePrice: {
          amount: 0,
          currency: "USD",
        },
        discountAllowed: false,
        specificChoices: [
          {
            type: "PRODUCT",
            nameOverride: "string",
            pricingModel: "string",
            overridePrice: {
              amount: 0,
              currency: "USD",
            },
            discountAllowed: false,
            asset: {
              type: "string",
              url: "string",
              altText: "string",
              title: "string",
              tags: ["string"],
            },
          },
        ],
      },
    },
  ],
  primaryAsset: {
    type: "string",
    url: "string",
    altText: "string",
    title: "string",
    tags: ["string"],
    id: "string",
  },
  primaryCategory: {
    id: "string",
  },
  brand: {
    brandName: "string",
  },
};
