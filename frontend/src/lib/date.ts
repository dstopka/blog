import { DateTime } from 'luxon';

export function formatDate(dateStr: string) {
    var date = DateTime.fromISO(dateStr);
    return date.toLocaleString(DateTime.DATE_MED, { locale: 'en-us' });
};