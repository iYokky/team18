/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCheckIn,
    EntCheckInFromJSON,
    EntCheckInFromJSONTyped,
    EntCheckInToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusCheckInEdges
 */
export interface EntStatusCheckInEdges {
    /**
     * Checkins holds the value of the checkins edge.
     * @type {Array<EntCheckIn>}
     * @memberof EntStatusCheckInEdges
     */
    checkins?: Array<EntCheckIn>;
}

export function EntStatusCheckInEdgesFromJSON(json: any): EntStatusCheckInEdges {
    return EntStatusCheckInEdgesFromJSONTyped(json, false);
}

export function EntStatusCheckInEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusCheckInEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkins': !exists(json, 'checkins') ? undefined : ((json['checkins'] as Array<any>).map(EntCheckInFromJSON)),
    };
}

export function EntStatusCheckInEdgesToJSON(value?: EntStatusCheckInEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkins': value.checkins === undefined ? undefined : ((value.checkins as Array<any>).map(EntCheckInToJSON)),
    };
}

