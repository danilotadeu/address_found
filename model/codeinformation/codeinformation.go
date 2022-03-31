package codeinformation

import "encoding/xml"

//CodeInformationRequest is a struct to request
type CodeInformationRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password,omitempty"`
	Customer struct {
		ID       string `json:"id"`
		SyncFlag bool   `json:"syncFlag"`
	} `json:"customer"`
}

type ResponseData struct {
	Err  *CodeInformationResponseProviderError `json:"err"`
	Data *CodeInformationServiceResponse       `json:"data"`
}

//CodeInformationResponse is a response CodeInformation
type CodeInformationResponse struct {
	CustomerCode string `json:"customerCode"`
}

type CodeInformationResponseError struct {
	Description string   `json:"description"`
	Provider    Provider `json:"provider"`
}

type CodeInformationResponseProviderError struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata"`
		Fault struct {
			Text      string `xml:",chardata"`
			Faultcode struct {
				Text string `xml:",chardata"`
				Ns0  string `xml:"ns0,attr"`
			} `xml:"faultcode"`
			Faultstring struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"faultstring"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type Provider struct {
	ServiceName  string `json:"serviceName"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

//CustomerRetrieveRequest is a request XML of provider r-customer-code-information-provider
type CustomerRetrieveRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Cus     string   `xml:"cus,attr"`
	Header  struct {
		Text     string `xml:",chardata"`
		Security struct {
			Text           string `xml:",chardata"`
			Actor          string `xml:"actor,attr"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
			Wsse           string `xml:"wsse,attr"`
			UsernameToken  struct {
				Text     string `xml:",chardata"`
				Wsu      string `xml:"wsu,attr"`
				Username string `xml:"Username"`
				Password struct {
					Text string `xml:",chardata"`
					Type string `xml:"Type,attr"`
				} `xml:"Password"`
			} `xml:"UsernameToken"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text                    string `xml:",chardata"`
		CustomerRetrieveRequest struct {
			Text            string `xml:",chardata"`
			InputAttributes struct {
				Text         string `xml:",chardata"`
				CustomerRead struct {
					Text       string `xml:",chardata"`
					CsId       int    `xml:"csId"`
					CsIdPub    string `xml:"csIdPub"`
					SyncWithDb bool   `xml:"syncWithDb"`
				} `xml:"customerRead"`
				PaymentArrangementsRead struct {
					Text        string `xml:",chardata"`
					FlagCurrent string `xml:"flagCurrent"`
					CsId        int    `xml:"csId"`
					CsIdPub     string `xml:"csIdPub"`
				} `xml:"paymentArrangementsRead"`
				AddressesRead struct {
					Text    string `xml:",chardata"`
					CsId    int    `xml:"csId"`
					CsIdPub string `xml:"csIdPub"`
				} `xml:"addressesRead"`
				CustomerInfoRead struct {
					Text    string `xml:",chardata"`
					CsId    int    `xml:"csId"`
					CsIdPub string `xml:"csIdPub"`
				} `xml:"customerInfoRead"`
			} `xml:"inputAttributes"`
			SessionChangeRequest struct {
				Text   string `xml:",chardata"`
				Values struct {
					Text string `xml:",chardata"`
					Item struct {
						Text  string `xml:",chardata"`
						Key   string `xml:"key"`
						Value string `xml:"value"`
					} `xml:"item"`
				} `xml:"values"`
			} `xml:"sessionChangeRequest"`
		} `xml:"customerRetrieveRequest"`
	} `xml:"Body"`
}

type CodeInformationServiceResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Cus     string   `xml:"cus,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                     string `xml:",chardata"`
		CustomerRetrieveResponse struct {
			Text         string `xml:",chardata"`
			CustomerRead struct {
				Text     string `xml:",chardata"`
				TbAmount struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"tbAmount"`
				CreatedByUser   string `xml:"createdByUser"`
				TbMode          string `xml:"tbMode"`
				TbPurpose       string `xml:"tbPurpose"`
				SioActionType   string `xml:"sioActionType"`
				SioThresholdAmt struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"sioThresholdAmt"`
				MaxCarryOverLenght           string `xml:"maxCarryOverLenght"`
				MaxCarryOverPeriodType       string `xml:"maxCarryOverPeriodType"`
				IsIndividualOverdiscDisabled string `xml:"isIndividualOverdiscDisabled"`
				FamilyId                     string `xml:"familyId"`
				ListOfAssignments            struct {
					Text string `xml:",chardata"`
					Item struct {
						Text             string `xml:",chardata"`
						FamilyId         string `xml:"familyId"`
						CompletionStatus string `xml:"completionStatus"`
						ValidFrom        string `xml:"validFrom"`
					} `xml:"item"`
				} `xml:"listOfAssignments"`
				CsId                     string `xml:"csId"`
				CsIdPub                  string `xml:"csIdPub"`
				CsCode                   string `xml:"csCode"`
				CsStatus                 string `xml:"csStatus"`
				CsStatusDate             string `xml:"csStatusDate"`
				CsActivationDate         string `xml:"csActivationDate"`
				CsIdHigh                 string `xml:"csIdHigh"`
				CsIdHighPub              string `xml:"csIdHighPub"`
				CsLevelCode              string `xml:"csLevelCode"`
				PaymentResp              string `xml:"paymentResp"`
				PrgCode                  string `xml:"prgCode"`
				Rpcode                   string `xml:"rpcode"`
				RpcodePub                string `xml:"rpcodePub"`
				TradeCode                string `xml:"tradeCode"`
				AreaCode                 string `xml:"areaCode"`
				CostId                   string `xml:"costId"`
				CostCodePub              string `xml:"costCodePub"`
				CsPassword               string `xml:"csPassword"`
				RsCode                   string `xml:"rsCode"`
				WpCode                   string `xml:"wpCode"`
				SrCode                   string `xml:"srCode"`
				CsRemark1                string `xml:"csRemark1"`
				CsRemark2                string `xml:"csRemark2"`
				CsBillInformation        string `xml:"csBillInformation"`
				ExpectPayCurrencyId      string `xml:"expectPayCurrencyId"`
				ExpectPayCurrencyIdPub   string `xml:"expectPayCurrencyIdPub"`
				CsConvratetypePayment    string `xml:"csConvratetypePayment"`
				CsConvratetypePaymentPub string `xml:"csConvratetypePaymentPub"`
				RefundCurrencyId         string `xml:"refundCurrencyId"`
				RefundCurrencyIdPub      string `xml:"refundCurrencyIdPub"`
				CsConvratetypeRefund     string `xml:"csConvratetypeRefund"`
				CsConvratetypeRefundPub  string `xml:"csConvratetypeRefundPub"`
				CsCrcheckAgreed          string `xml:"csCrcheckAgreed"`
				CustcatCode              string `xml:"custcatCode"`
				CustcatCodePub           string `xml:"custcatCodePub"`
				CsJurisdictionId         string `xml:"csJurisdictionId"`
				CsJurisdictionCode       string `xml:"csJurisdictionCode"`
				CsIncorporatedInd        string `xml:"csIncorporatedInd"`
				CsBillcycle              string `xml:"csBillcycle"`
				CsBillcycleDesc          string `xml:"csBillcycleDesc"`
				CsLimitTr1               string `xml:"csLimitTr1"`
				CsLimitTr1Pub            string `xml:"csLimitTr1Pub"`
				CsLimitTr2               string `xml:"csLimitTr2"`
				CsLimitTr2Pub            string `xml:"csLimitTr2Pub"`
				CsLimitTr3               string `xml:"csLimitTr3"`
				CsLimitTr3Pub            string `xml:"csLimitTr3Pub"`
				CsClimit                 struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"csClimit"`
				CsContresp   string `xml:"csContresp"`
				CsLastBcDate string `xml:"csLastBcDate"`
				CsDeposit    struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"csDeposit"`
				CsCreditscore    string `xml:"csCreditscore"`
				CsCscreditStatus string `xml:"csCscreditStatus"`
				CsCscreditRemark string `xml:"csCscreditRemark"`
				CsCscreditDate   string `xml:"csCscreditDate"`
				CsDunning        string `xml:"csDunning"`
				CsDunningMode    string `xml:"csDunningMode"`
				CsPrepayment     string `xml:"csPrepayment"`
				CsCollector      string `xml:"csCollector"`
				CsFcId           string `xml:"csFcId"`
				CsFcIdPub        string `xml:"csFcIdPub"`
				CsCscurbalance   struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"csCscurbalance"`
				CsPrevbalance struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"csPrevbalance"`
				CsUnbilledAmount struct {
					Text     string `xml:",chardata"`
					Amount   string `xml:"amount"`
					Currency string `xml:"currency"`
				} `xml:"csUnbilledAmount"`
				CsPaybehaviour          string `xml:"csPaybehaviour"`
				CsNoBouncePay           string `xml:"csNoBouncePay"`
				CsDealerid              string `xml:"csDealerid"`
				CsDealeridPub           string `xml:"csDealeridPub"`
				CsInitPrepaidContrOwner string `xml:"csInitPrepaidContrOwner"`
				CsCreationDate          string `xml:"csCreationDate"`
				CsDeactivationDate      string `xml:"csDeactivationDate"`
				CsSuspensionDate        string `xml:"csSuspensionDate"`
				CsReactivationDate      string `xml:"csReactivationDate"`
				LastModificationUser    string `xml:"lastModificationUser"`
				PaymentRespId           string `xml:"paymentRespId"`
				CsPaymntRespCode        string `xml:"csPaymntRespCode"`
				PartyRoleAssignments    struct {
					Text string `xml:",chardata"`
					Item struct {
						Text                 string `xml:",chardata"`
						PartyRoleId          string `xml:"partyRoleId"`
						PartyRoleShname      string `xml:"partyRoleShname"`
						PartyRoleName        string `xml:"partyRoleName"`
						PartyRoleDescription string `xml:"partyRoleDescription"`
					} `xml:"item"`
				} `xml:"partyRoleAssignments"`
				PartyType             string `xml:"partyType"`
				CollInd               string `xml:"collInd"`
				MarkedForRerating     string `xml:"markedForRerating"`
				ExternalCustomerId    string `xml:"externalCustomerId"`
				ExternalCustomerSetId string `xml:"externalCustomerSetId"`
				MaxCsLevel            string `xml:"maxCsLevel"`
				CreditScore           string `xml:"creditScore"`
				CreditScoreProvider   string `xml:"creditScoreProvider"`
				CsCscreditUser        string `xml:"csCscreditUser"`
				PerformCreditScoring  string `xml:"performCreditScoring"`
				Cslvlname             string `xml:"cslvlname"`
				AdrNationality        string `xml:"adrNationality"`
				AdrNationalityPub     string `xml:"adrNationalityPub"`
				AdrEmployee           string `xml:"adrEmployee"`
				CsTestbillrun         string `xml:"csTestbillrun"`
			} `xml:"customerRead"`
			PaymentArrangementsRead struct {
				Text                      string `xml:",chardata"`
				ListOfPaymentArrangements struct {
					Text string `xml:",chardata"`
					Item struct {
						Text           string `xml:",chardata"`
						CspId          string `xml:"cspId"`
						CspSeqno       string `xml:"cspSeqno"`
						CspPmntId      string `xml:"cspPmntId"`
						CspPmntIdPub   string `xml:"cspPmntIdPub"`
						CspAccno       string `xml:"cspAccno"`
						CspAccowner    string `xml:"cspAccowner"`
						CspBankcode    string `xml:"cspBankcode"`
						CspBankname    string `xml:"cspBankname"`
						CspBankzip     string `xml:"cspBankzip"`
						CspBankcity    string `xml:"cspBankcity"`
						CspBankstreet  string `xml:"cspBankstreet"`
						CspCcagCode    string `xml:"cspCcagCode"`
						CspCcagCodePub string `xml:"cspCcagCodePub"`
						CspCcvaliddt   string `xml:"cspCcvaliddt"`
						CspCeiling     struct {
							Text     string `xml:",chardata"`
							Amount   string `xml:"amount"`
							Currency string `xml:"currency"`
						} `xml:"cspCeiling"`
						CspBankstate      string `xml:"cspBankstate"`
						CspBankcounty     string `xml:"cspBankcounty"`
						CspOrdernumber    string `xml:"cspOrdernumber"`
						CspBankstreetno   string `xml:"cspBankstreetno"`
						CspBankcountry    string `xml:"cspBankcountry"`
						CspBankcountryPub string `xml:"cspBankcountryPub"`
						CspSwiftcode      string `xml:"cspSwiftcode"`
						CspActUsed        string `xml:"cspActUsed"`
						CspValidFrom      string `xml:"cspValidFrom"`
						AuthOk            string `xml:"authOk"`
						AuthDate          string `xml:"authDate"`
						AuthNo            string `xml:"authNo"`
						AuthCredit        struct {
							Text     string `xml:",chardata"`
							Amount   string `xml:"amount"`
							Currency string `xml:"currency"`
						} `xml:"authCredit"`
						AuthTn     string `xml:"authTn"`
						AuthRemark string `xml:"authRemark"`
						CsId       string `xml:"csId"`
						CsIdPub    string `xml:"csIdPub"`
					} `xml:"item"`
				} `xml:"listOfPaymentArrangements"`
			} `xml:"paymentArrangementsRead"`
			AddressesRead struct {
				Text               string `xml:",chardata"`
				ListOfAllAddresses struct {
					Text string `xml:",chardata"`
					Item struct {
						Text                  string `xml:",chardata"`
						AdrSeq                string `xml:"adrSeq"`
						AdrRoles              string `xml:"adrRoles"`
						AdrJurTaxOverridden   string `xml:"adrJurTaxOverridden"`
						AdrTempbillOverridden string `xml:"adrTempbillOverridden"`
						AdrDeleted            string `xml:"adrDeleted"`
						TtlId                 string `xml:"ttlId"`
						TtlIdPub              string `xml:"ttlIdPub"`
						AdrLname              string `xml:"adrLname"`
						AdrName               string `xml:"adrName"`
						AdrFname              string `xml:"adrFname"`
						AdrStreet             string `xml:"adrStreet"`
						AdrStreetno           string `xml:"adrStreetno"`
						AdrZip                string `xml:"adrZip"`
						AdrCity               string `xml:"adrCity"`
						CountryId             string `xml:"countryId"`
						CountryIdPub          string `xml:"countryIdPub"`
						LngCode               string `xml:"lngCode"`
						LngCodePub            string `xml:"lngCodePub"`
						AdrNote1              string `xml:"adrNote1"`
						AdrNote2              string `xml:"adrNote2"`
						AdrNote3              string `xml:"adrNote3"`
						AdrJbdes              string `xml:"adrJbdes"`
						AdrPhn1Area           string `xml:"adrPhn1Area"`
						AdrPhn1               string `xml:"adrPhn1"`
						AdrPhn2Area           string `xml:"adrPhn2Area"`
						AdrPhn2               string `xml:"adrPhn2"`
						AdrFaxArea            string `xml:"adrFaxArea"`
						AdrFax                string `xml:"adrFax"`
						AdrMname              string `xml:"adrMname"`
						AdrEmail              string `xml:"adrEmail"`
						AdrSmsno              string `xml:"adrSmsno"`
						AdrYears              string `xml:"adrYears"`
						AdrState              string `xml:"adrState"`
						AdrCounty             string `xml:"adrCounty"`
						AdrValiddate          string `xml:"adrValiddate"`
						AdrInccode            string `xml:"adrInccode"`
						AdrUrgent             string `xml:"adrUrgent"`
						AdrForward            string `xml:"adrForward"`
						AdrLocation1          string `xml:"adrLocation1"`
						AdrLocation2          string `xml:"adrLocation2"`
						AdrRemark             string `xml:"adrRemark"`
						AdrTaxno              string `xml:"adrTaxno"`
						AdrCompno             string `xml:"adrCompno"`
						IdtypeCode            string `xml:"idtypeCode"`
						AdrIdno               string `xml:"adrIdno"`
						AdrBirthdt            string `xml:"adrBirthdt"`
						AdrSocialseno         string `xml:"adrSocialseno"`
						AdrDrivelicence       string `xml:"adrDrivelicence"`
						AdrSex                string `xml:"adrSex"`
						AdrEmployer           string `xml:"adrEmployer"`
						AdrCusttype           string `xml:"adrCusttype"`
						MasCode               string `xml:"masCode"`
						MasCodePub            string `xml:"masCodePub"`
						AdrNationality        string `xml:"adrNationality"`
						AdrNationalityPub     string `xml:"adrNationalityPub"`
						AdrEmployee           string `xml:"adrEmployee"`
					} `xml:"item"`
				} `xml:"listOfAllAddresses"`
				CsId    string `xml:"csId"`
				CsIdPub string `xml:"csIdPub"`
			} `xml:"addressesRead"`
			CustomerInfoRead struct {
				Text    string `xml:",chardata"`
				Text01  string `xml:"text01"`
				Text02  string `xml:"text02"`
				Text03  string `xml:"text03"`
				Text04  string `xml:"text04"`
				Text05  string `xml:"text05"`
				Text06  string `xml:"text06"`
				Text07  string `xml:"text07"`
				Text08  string `xml:"text08"`
				Text09  string `xml:"text09"`
				Text10  string `xml:"text10"`
				Text11  string `xml:"text11"`
				Text12  string `xml:"text12"`
				Text13  string `xml:"text13"`
				Text14  string `xml:"text14"`
				Text15  string `xml:"text15"`
				Text16  string `xml:"text16"`
				Text17  string `xml:"text17"`
				Text18  string `xml:"text18"`
				Text19  string `xml:"text19"`
				Text20  string `xml:"text20"`
				Text21  string `xml:"text21"`
				Text22  string `xml:"text22"`
				Text23  string `xml:"text23"`
				Text24  string `xml:"text24"`
				Text25  string `xml:"text25"`
				Text26  string `xml:"text26"`
				Text27  string `xml:"text27"`
				Text28  string `xml:"text28"`
				Text29  string `xml:"text29"`
				Text30  string `xml:"text30"`
				Check01 string `xml:"check01"`
				Check02 string `xml:"check02"`
				Check03 string `xml:"check03"`
				Check04 string `xml:"check04"`
				Check05 string `xml:"check05"`
				Check06 string `xml:"check06"`
				Check07 string `xml:"check07"`
				Check08 string `xml:"check08"`
				Check09 string `xml:"check09"`
				Check10 string `xml:"check10"`
				Check11 string `xml:"check11"`
				Check12 string `xml:"check12"`
				Check13 string `xml:"check13"`
				Check14 string `xml:"check14"`
				Check15 string `xml:"check15"`
				Check16 string `xml:"check16"`
				Check17 string `xml:"check17"`
				Check18 string `xml:"check18"`
				Check19 string `xml:"check19"`
				Check20 string `xml:"check20"`
				Combo01 string `xml:"combo01"`
				Combo02 string `xml:"combo02"`
				Combo03 string `xml:"combo03"`
				Combo04 string `xml:"combo04"`
				Combo05 string `xml:"combo05"`
				Combo06 string `xml:"combo06"`
				Combo07 string `xml:"combo07"`
				Combo08 string `xml:"combo08"`
				Combo09 string `xml:"combo09"`
				Combo10 string `xml:"combo10"`
				Combo11 string `xml:"combo11"`
				Combo12 string `xml:"combo12"`
				Combo13 string `xml:"combo13"`
				Combo14 string `xml:"combo14"`
				Combo15 string `xml:"combo15"`
				Combo16 string `xml:"combo16"`
				Combo17 string `xml:"combo17"`
				Combo18 string `xml:"combo18"`
				Combo19 string `xml:"combo19"`
				Combo20 string `xml:"combo20"`
				CsId    string `xml:"csId"`
				CsIdPub string `xml:"csIdPub"`
			} `xml:"customerInfoRead"`
		} `xml:"customerRetrieveResponse"`
	} `xml:"Body"`
}

func (customer *CodeInformationServiceResponse) Transform() string {
	return customer.Body.CustomerRetrieveResponse.CustomerRead.CsCode
}
