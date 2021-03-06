package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//ErrorResponse structure for to accept Error Code and Error Message
type ErrorResponse struct {
	Code    int
	Message string
}

//SuccessResponse structure to accept Success Response Code and message
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

//Claims mean structure for name email from payload
type Claims struct {
	Name   string
	Email  string
	Userid int
	jwt.StandardClaims
}

//RegistrationParams structure for name email and password
type RegistrationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//LoginParams structure for email and password for the login form request
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//SuccessfulLoginResponse structure for Name, email and Token
type SuccessfulLoginResponse struct {
	Name      string
	Email     string
	AuthToken string
	Userid    int `json:"user_id" bson:"user_id"`
}

//UserDetails structure for user detail
type UserDetails struct {
	Name     string
	Email    string
	Password string
	Userid   int `json:"user_id" bson:"user_id"`
}

//CompleteUserDetails is a completed version of user detail
type CompleteUserDetails struct {
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
	Userid int    `json:"user_id" bson:"user_id"`
}

//MasterPart a structure for Master Part Table
type MasterPart struct {
	Partid       string    `json:"part_id" bson:"part_id"`
	Mgcatid      string    `json:"mg_cat_id" bson:"mg_cat_id"`
	Mgcatname    string    `json:"mg_cat_name" bson:"mg_cat_name"`
	Partcode     string    `json:"part_code" bson:"part_code"`
	Partname     string    `json:"part_name" bson:"part_name"`
	Partunit     string    `json:"part_unit" bson:"part_unit"`
	Supplierid   string    `json:"supplier_id" bson:"supplier_id"`
	Suppliername string    `json:"supplier_name" bson:"supplier_name"`
	Minstock     string    `json:"min_stock" bson:"min_stock"`
	Costprice    string    `json:"cost_price" bson:"cost_price"`
	Expireddate  string    `json:"expired_date" bson:"expired_date"`
	Siteid       string    `json:"site_id" bson:"site_id"`
	Sitename     string    `json:"site_name" bson:"site_name"`
	Partnotes    string    `json:"part_notes" bson:"part_notes"`
	Userid       int       `json:"user_id" bson:"user_id"`
	Username     string    `json:"name" bson:"name"`
	Stock        int       `json:"quantity" bson:"quantity"`
	Insertdate   time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate   time.Time `json:"update_date" bson:"update_date"`
}

//MasterGoods a structure for Master Finished Goods Table
type MasterGoods struct {
	Fgid            string    `json:"fg_id" bson:"fg_id"`
	Mgcatid         string    `json:"mg_cat_id" bson:"mg_cat_id"`
	Mgcatname       string    `json:"mg_cat_name" bson:"mg_cat_name"`
	Fgcode          string    `json:"fg_code" bson:"fg_code"`
	Fgname          string    `json:"fg_name" bson:"fg_name"`
	Fgunit          string    `json:"fg_unit" bson:"fg_unit"`
	Minstock        string    `json:"min_stock" bson:"min_stock"`
	Costprice       string    `json:"production_cost" bson:"production_cost"`
	Percentmarkup   string    `json:"percent_markup" bson:"percent_markup"`
	Percentdiscount string    `json:"percent_discount" bson:"percent_discount"`
	Netprice        string    `json:"net_price" bson:"net_price"`
	Expireddate     string    `json:"expired_date" bson:"expired_date"`
	Siteid          string    `json:"site_id" bson:"site_id"`
	Fgnotes         string    `json:"fg_notes" bson:"fg_notes"`
	Userid          int       `json:"user_id" bson:"user_id"`
	Username        string    `json:"name" bson:"name"`
	Insertdate      time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate      time.Time `json:"update_date" bson:"update_date"`
}

//AgregateMasterPart Pipeline from agregation data master part
type AgregateMasterPart struct {
	Partid       string    `json:"part.part_id" bson:"part.part_id"`
	Mgcatid      string    `json:"part.mg_cat_id" bson:"part.mg_cat_id"`
	Mgcatname    string    `json:"category.mg_cat_name" bson:"category.mg_cat_name"`
	Partcode     string    `json:"part.part_code" bson:"part.part_code"`
	Partname     string    `json:"part_name" bson:"part_name"`
	Partunit     string    `json:"part_unit" bson:"part_unit"`
	Supplierid   string    `json:"supplier_id" bson:"supplier_id"`
	Suppliername string    `json:"supplier_name" bson:"supplier_name"`
	Minstock     string    `json:"min_stock" bson:"min_stock"`
	Costprice    string    `json:"cost_price" bson:"cost_price"`
	Expireddate  string    `json:"expired_date" bson:"expired_date"`
	Siteid       string    `json:"site_id" bson:"site_id"`
	Sitename     string    `json:"site_name" bson:"site_name"`
	Partnotes    string    `json:"part_notes" bson:"part_notes"`
	Userid       int       `json:"user_id" bson:"user_id"`
	Username     string    `json:"name" bson:"name"`
	Stock        int       `json:"quantity" bson:"quantity"`
	Insertdate   time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate   time.Time `json:"update_date" bson:"update_date"`
}

//AgregateMasterGoods a structure for Master Finished Goods Table
type AgregateMasterGoods struct {
	Fgid            string    `json:"good.fg_id" bson:"good.fg_id"`
	Mgcatid         string    `json:"good.mg_cat_id" bson:"good.mg_cat_id"`
	Mgcatname       string    `json:"category.mg_cat_name" bson:"category.mg_cat_name"`
	Fgcode          string    `json:"good.fg_code" bson:"good.fg_code"`
	Fgname          string    `json:"fg_name" bson:"fg_name"`
	Fgunit          string    `json:"fg_unit" bson:"fg_unit"`
	Minstock        int       `json:"min_stock" bson:"min_stock"`
	Costprice       int       `json:"production_cost" bson:"production_cost"`
	Percentmarkup   int       `json:"percent_markup" bson:"percent_markup"`
	Percentdiscount int       `json:"percent_discount" bson:"percent_discount"`
	Netprice        int       `json:"net_price" bson:"net_price"`
	Expireddate     time.Time `json:"expired_date" bson:"expired_date"`
	Siteid          string    `json:"site_id" bson:"site_id"`
	Sitename        string    `json:"site_name" bson:"site_name"`
	Fgnotes         string    `json:"fg_notes" bson:"fg_notes"`
	Userid          string    `json:"user_id" bson:"user_id"`
	Username        string    `json:"name" bson:"name"`
	Stock           int       `json:"quantity" bson:"quantity"`
	Insertdate      time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate      time.Time `json:"update_date" bson:"update_date"`
}

//MstSupplier is a structure for Master Supplier
type MstSupplier struct {
	Supplierid    string `json:"supplier_id" bson:"supplier_id"`
	Suppliername  string `json:"supplier_name" bson:"supplier_name"`
	Supplieraddr  string `json:"supplier_addr" bson:"supplier_addr"`
	Supplieremail string `json:"supplier_email" bson:"supplier_email"`
	Supplierphone string `json:"supplier_phone" bson:"supplier_phone"`
}

//MstSite is a structure for Master Site
type MstSite struct {
	Siteid   string `json:"site_id" bson:"site_id"`
	Sitename string `json:"site_name" bson:"site_name"`
	Siteaddr string `json:"site_addr" bson:"site_addr"`
	Sitelong string `json:"site_long" bson:"site_long"`
	Sitelat  string `json:"site_lat" bson:"site_lat"`
}

//MstWarehouse is a structure for Master Warehouse
type MstWarehouse struct {
	Warehouseid   string `json:"wh_id" bson:"wh_id"`
	Warehousename string `json:"wh_name" bson:"wh_name"`
}

//TCategory is a structure for table category
type TCategory struct {
	Mgcatid   string `json:"mg_cat_id" bson:"mg_cat_id"`
	Mgcatname string `json:"mg_cat_name" bson:"mg_cat_name"`
}

//TDoc is a structure for Table Doc
type TDoc struct {
	Docnumber  int       `json:"doc_number" bson:"doc_number"`
	Doccatid   int       `json:"doc_cat_id" bson:"doc_cat_id"`
	Filename   string    `json:"file_name" bson:"file_name"`
	Userid     string    `json:"user_id" bson:"user_id"`
	Insertdate time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate time.Time `json:"update_date" bson:"update_date"`
}

//TPurchase is a structure for Table Purchase
type TPurchase struct {
	Purchaseid    string    `json:"purchase_id" bson:"purchase_id"`
	Supplierid    string    `json:"supplier_id" bson:"supplier_id"`
	Whid          string    `json:"wh_id" bson:"wh_id"`
	Partid        string    `json:"part_id" bson:"part_id"`
	Qty           string    `json:"qty" bson:"qty"`
	Purchasedate  time.Time `json:"purchase_date" bson:"purchase_date"`
	Estimatedcost string    `json:"estimated_cost" bson:"estimated_cost"`
	Invoice       string    `json:"invoice" bson:"invoice"`
	Receipt       string    `json:"receipt" bson:"receipt"`
	Buyerid       int       `json:"buyer_id" bson:"buyer_id"`
	Originatorid  int       `json:"originator_id" bson:"originator_id"`
	Userid        int       `json:"user_id" bson:"user_id"`
	Notes         string    `json:"notes" bson:"notes"`
	Statusflag    string    `json:"status_flag" bson:"status_flag"`
	Bidoutdate    time.Time `json:"bidout_date" bson:"bidout_date"`
	Closeddate    time.Time `json:"closed_date" bson:"closed_date"`
	Insertdate    time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate    time.Time `json:"update_date" bson:"update_date"`
}

//TPurchaseView is a structure for Table Purchase View
type TPurchaseView struct {
	Purchaseid     string    `json:"purchase_id" bson:"purchase_id"`
	Supplierid     string    `json:"purchase.supplier_id" bson:"purchase.supplier_id"`
	Suppliername   string    `json:"supplier.supplier_name" bson:"supplier.supplier_name"`
	Whid           string    `json:"warehouse.wh_id" bson:"warehouse.wh_id"`
	Whname         string    `json:"warehouse.wh_name" bson:"warehouse.wh_name"`
	Partid         string    `json:"purchase.part_id" bson:"purchase.part_id"`
	Partname       string    `json:"part.part_name" bson:"part.part_name"`
	Qty            int       `json:"purchase.qty" bson:"purchase.qty"`
	Stock          int       `json:"stock.quantity" bson:"stock.quantity"`
	Purchasedate   time.Time `json:"purchase.purchase_date" bson:"purchase.purchase_date"`
	Estimatedcost  int       `json:"supplier.estimated_cost" bson:"supplier.estimated_cost"`
	Invoiceid      string    `json:"purchase.invoice" bson:"purchase.invoice"`
	Invoicefile    string    `json:"docinvoice.file_name	" bson:"docinvoice.file_name"`
	Receiptid      string    `json:"purchase.receipt	" bson:"purchase.receipt"`
	Receiptfile    string    `json:"docreceipt.file_name	" bson:"docreceipt.file_name"`
	Buyerid        string    `json:"purchase.buyer_id	" bson:"purchase.buyer_id"`
	Buyername      string    `json:"buyer.name	" bson:"buyer.name"`
	Originatorid   string    `json:"purchase.originator_id	" bson:"purchase.originator_id"`
	Originatorname string    `json:"originator.name	" bson:"originator.name"`
	Userid         string    `json:"purchase.user_id" bson:"purchase.user_id"`
	Username       string    `json:"user.name" bson:"user.name"`
	Notes          string    `json:"purchase.notes" bson:"purchase.notes"`
	Statusflag     string    `json:"purchase.status_flag" bson:"purchase.status_flag"`
	StatusName     string    `json:"status.status_name" bson:"status.status_name"`
	Bidoutdate     time.Time `json:"purchase.bidout_date" bson:"purchase.bidout_date"`
	Closeddate     time.Time `json:"purchase.closed_date" bson:"purchase.closed_date"`
	Insertdate     time.Time `json:"purchase.insert_date" bson:"purchase.insert_date"`
	Updatedate     time.Time `json:"purchase.update_date" bson:"purchase.update_date"`
}

//TPurchaseViewTest is a structure for Table Purchase View
type TPurchaseViewTest struct {
	TPurchasePart struct {
	} `json:"part" bson:"part"`
}

//TPurchasePart is a structure for Table Purchasepart View
type TPurchasePart struct {
	Partname string `json:"part_name" bson:"part_name"`
}

//TSales is a struct for Table Sales
type TSales struct {
	Salesid    string    `json:"sales_id" bson:"sales_id"`
	Qtytotal   int       `json:"qty_total" bson:"qty_total"`
	Total      int       `json:"total" bson:"total"`
	Orderdate  time.Time `json:"order_date" bson:"order_date"`
	Invoice    string    `json:"invoice" bson:"invoice"`
	Receipt    string    `json:"receipt" bson:"receipt"`
	Userid     string    `json:"user_id" bson:"user_id"`
	Statusflag string    `json:"status_flag" bson:"status_flag"`
	Insertdate time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate time.Time `json:"update_date" bson:"update_date"`
}

//TSelesDetail is a structure for Table Sales Detail
type TSelesDetail struct {
	Sdid    string `json:"sd_id" bson:"sd_id"`
	Salesid string `json:"sales_id" bson:"sales_id"`
	Fgid    string `json:"fg_id" bson:"fg_id"`
	Qty     int    `json:"qty" bson:"qty"`
	Price   int    `json:"price" bson:"price"`
	Notes   string `json:"notes" bson:"notes"`
}

//TStock is a structure for Table Stock
type TStock struct {
	Stockid    string    `json:"stock_id" bson:"stock_id"`
	Mgid       string    `json:"mg_id" bson:"mg_id"`
	Quantity   int       `json:"quantity" bson:"quantity"`
	Insertdate time.Time `json:"insertdate" bson:"insert_date"`
	Updatedate time.Time `json:"update_date" bson:"update_date"`
}

//TMutation is a structure for Table Mutation
type TMutation struct {
	Mutationid    int       `json:"mutation_id" bson:"mutation_id"`
	Mutationtitle string    `json:"mutation_title" bson:"mutation_title"`
	Mutationnote  string    `json:"mutation_note" bson:"mutation_note"`
	Docnumber     int       `json:"doc_number" bson:"doc_number"`
	Whidfrom      string    `json:"wh_idfrom" bson:"wh_id_from"`
	Whidto        string    `json:"wh_id_to" bson:"wh_id_to"`
	Siteidfrom    string    `json:"site_id_from" bson:"site_id_from"`
	Siteidto      string    `json:"site_id_to" bson:"site_id_to"`
	Mgid          string    `json:"mg_id" bson:"mg_id"`
	Mgidto        string    `json:"mg_id_to" bson:"mg_id_to"`
	Qty           int       `json:"qty" bson:"qty"`
	Userid        string    `json:"user_id" bson:"user_id"`
	Insertdate    time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate    time.Time `json:"update_date" bson:"update_date"`
	Statusflag    time.Time `json:"status_flag" bson:"status_flag"`
}

//TUse is a structure for Table Use
type TUse struct {
	Useid      int       `json:"use_id" bson:"use_id"`
	Usetitle   string    `json:"use_title" bson:"use_title"`
	Usenote    string    `json:"use_note" bson:"use_note"`
	Docnumber  int       `json:"doc_number" bson:"doc_number"`
	Mgid       string    `json:"mg_id" bson:"mg_id"`
	Qty        int       `json:"qty" bson:"qty"`
	Userid     string    `json:"user_id" bson:"user_id"`
	Siteid     string    `json:"site_id" bson:"site_id"`
	Insertdate time.Time `json:"insert_date" bson:"insert_date"`
	Updatedate time.Time `json:"update_date" bson:"update_date"`
	Statusflag string    `json:"status_flag" bson:"status_flag"`
}
