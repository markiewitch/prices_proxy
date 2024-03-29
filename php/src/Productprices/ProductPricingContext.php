<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pricesproxy/prices_proxy.proto

namespace Productprices;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>productprices.ProductPricingContext</code>
 */
class ProductPricingContext extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 productId = 1;</code>
     */
    private $productId = 0;
    /**
     * Generated from protobuf field <code>.productprices.PricingContext context = 2;</code>
     */
    private $context = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $productId
     *     @type \Productprices\PricingContext $context
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Pricesproxy\PricesProxy::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 productId = 1;</code>
     * @return int
     */
    public function getProductId()
    {
        return $this->productId;
    }

    /**
     * Generated from protobuf field <code>int32 productId = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setProductId($var)
    {
        GPBUtil::checkInt32($var);
        $this->productId = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.productprices.PricingContext context = 2;</code>
     * @return \Productprices\PricingContext
     */
    public function getContext()
    {
        return $this->context;
    }

    /**
     * Generated from protobuf field <code>.productprices.PricingContext context = 2;</code>
     * @param \Productprices\PricingContext $var
     * @return $this
     */
    public function setContext($var)
    {
        GPBUtil::checkMessage($var, \Productprices\PricingContext::class);
        $this->context = $var;

        return $this;
    }

}

